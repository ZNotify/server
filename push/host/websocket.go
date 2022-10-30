package host

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/db/entity"
	"notify-api/db/model"
	pushTypes "notify-api/push/types"
	"notify-api/serve/types"
	"notify-api/user"
	"notify-api/utils"
)

const (
	writeWait = 10 * time.Second

	timeout = 30 * time.Second

	pingPeriod = (timeout * 7) / 10

	maxMessageSize = 512
)

type Client struct {
	manager *Manager

	conn *websocket.Conn

	send *utils.UnboundedChan[*entity.Message]

	userID   string
	deviceID string

	once sync.Once
}

type WebSocketHost struct {
	manager *Manager
}

type Manager struct {
	userClients map[string]map[*Client]bool

	register chan *Client

	unregister chan *Client

	broadcast chan *entity.Message
}

type wsMessage entity.Message

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *Client) writeRoutine() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send.Out:
			if !ok {
				_ = c.conn.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(writeWait))
				return
			}

			wsMsg := wsMessage(*msg)

			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.conn.WriteJSON(wsMsg)
			if err != nil {
				zap.S().Infof("write message error: %v", err)
				c.Close()
				return
			}

			err = model.TokenUtils.CreateOrUpdate(c.userID, c.deviceID, "WebSocketHost", msg.CreatedAt.Format(time.RFC3339Nano))
			if err != nil {
				zap.S().Errorf("create or update token error: %v", err)
				continue
			}
		case <-ticker.C:
			if err := c.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				zap.S().Infof("ping error: %v", err)
				c.Close()
				return
			}
		}
	}
}

func (c *Client) readRoutine() {
	defer func() {
		c.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	for {
		_, _, err := c.conn.NextReader()
		if err != nil {
			normalCodes := []int{websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseInternalServerErr}
			if websocket.IsUnexpectedCloseError(err, normalCodes...) {
				zap.S().Infof("client %s %s read routine exit %v", c.userID, c.deviceID[0:7], err)
			}
			break
		}
	}
}

func (c *Client) Close() {
	zap.S().Debugf("client %s %s try close", c.userID, c.deviceID[0:7])
	c.once.Do(func() {
		zap.S().Debugf("client %s %s real close", c.userID, c.deviceID[0:7])
		_ = c.conn.Close()
		c.manager.unregister <- c
		close(c.send.In)
	})
}

func (h *WebSocketHost) clientManageRoutine() {
	zap.S().Debug("client manage routine start")
	deleteClient := func(client *Client) {
		if userMap, ok := h.manager.userClients[client.userID]; ok {
			if _, ok := userMap[client]; ok {
				delete(userMap, client)
			}
		}
	}

	for {
		select {
		case client := <-h.manager.register:
			h.manager.userClients[client.userID][client] = true

		case client := <-h.manager.unregister:
			deleteClient(client)

		case msg := <-h.manager.broadcast:
			for client := range h.manager.userClients[msg.UserID] {
				select {
				case client.send.In <- msg:
				default:
					deleteClient(client)
				}
			}
		}
	}
}

func (h *WebSocketHost) Start() error {
	go h.clientManageRoutine()
	return nil
}

func (h *WebSocketHost) Init() error {
	h.manager = &Manager{
		userClients: make(map[string]map[*Client]bool),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		broadcast:   make(chan *entity.Message),
	}

	for _, v := range user.Controller.Users() {
		h.manager.userClients[v] = make(map[*Client]bool)
	}

	return nil
}

func (h *WebSocketHost) HandlerPath() string {
	return "/host/conn"
}

func (h *WebSocketHost) HandlerMethod() string {
	return "GET"
}

func (h *WebSocketHost) Handler(context *types.Ctx) {
	userID := context.UserID

	deviceId := context.GetHeader("X-Device-ID")
	if deviceId == "" {
		zap.S().Infof("user %s connect without device ID", userID)
		context.JSONError(http.StatusBadRequest, errors.New("no device id"))
		return
	}

	token, err := model.TokenUtils.GetDeviceToken(userID, deviceId)
	if err != nil {
		if err == model.ErrNotFound {
			zap.S().Infof("user %s device %s token not found", userID, deviceId)
			context.JSONError(http.StatusUnauthorized, errors.New("token not found"))
			return
		} else {
			zap.S().Errorf("get user %s token error: %v", userID, err)
			context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
			return
		}
	}
	if token.Channel != h.Name() {
		zap.S().Infof("user %s channel not match", userID)
		context.JSONError(http.StatusBadRequest, errors.New("user current channel is not WebSocket"))
		return
	}

	sinceTime, err := time.Parse(time.RFC3339Nano, token.Token)
	if err != nil {
		zap.S().Infof("parse time error: %v", err)
		context.JSONError(http.StatusBadRequest, errors.WithStack(err))
		return
	}
	// 2022-09-18T11:14:00+08:00 as zero point

	pendingMessages, err := model.MessageUtils.GetUserMessageAfter(userID, sinceTime)
	if err != nil {
		zap.S().Errorf("get user %s message error: %v", userID, err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		zap.S().Errorf("upgrade error: %v", err)
		return
	}
	client := &Client{
		manager:  h.manager,
		conn:     conn,
		send:     utils.NewUnboundedChan[*entity.Message](2),
		userID:   userID,
		deviceID: deviceId,
		once:     sync.Once{},
	}
	h.manager.register <- client

	go client.writeRoutine()
	go client.readRoutine()

	for _, msg := range pendingMessages {
		msg := msg
		client.send.In <- &msg
	}
}

func (h *WebSocketHost) Send(msg *pushTypes.Message) error {
	eMsg := &entity.Message{
		ID:        msg.ID,
		UserID:    msg.UserID,
		Title:     msg.Title,
		Content:   msg.Content,
		Long:      msg.Long,
		CreatedAt: msg.CreatedAt,
	}
	h.manager.broadcast <- eMsg
	return nil
}

func (h *WebSocketHost) Name() string {
	return "WebSocketHost"
}
