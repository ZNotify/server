package host

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"notify-api/db/entity"
	"notify-api/db/model"
	"notify-api/push/types"
	"notify-api/serve/middleware"
	"notify-api/user"
	"notify-api/utils"
	"sync"
	"time"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

type Client struct {
	manager *Manager

	conn *websocket.Conn

	send *utils.UnboundedChan[*entity.Message]

	userID string

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

func (c *Client) sendRoutine() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := c.conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		select {
		case msg, ok := <-c.send.Out:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			wsMsg := wsMessage(*msg)
			err := c.conn.WriteJSON(wsMsg)
			if err != nil {
				log.Printf("write message error: %v", err)
				c.Close()
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.Close()
				return
			}
		}
	}
}

func (c *Client) readRoutine() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			return
		}
	}()
	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { _ = c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			normalCodes := []int{websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure}
			if websocket.IsUnexpectedCloseError(err, normalCodes...) {
				log.Printf("websocket close error: %v", err)
			}
			break
		}
	}
}

func (c *Client) Close() {
	c.once.Do(func() {
		err := c.conn.Close()
		if err != nil {
			log.Printf("close client error: %v", err)
			return
		}
		c.manager.unregister <- c
		close(c.send.In)
	})
}

func (h *WebSocketHost) Handler(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)
	since := context.GetHeader("X-Message-Since")
	if since == "" {
		log.Printf("user %s connect without since header", userID)
		context.AbortWithStatus(http.StatusBadRequest)
	}
	// parse RFC3339
	sinceTime, err := time.Parse(time.RFC3339Nano, since)
	if err != nil {
		log.Printf("parse time error: %v", err)
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// 2022-09-18T11:14:00+08:00 as zero point

	pendingMessages, err := model.MessageUtils.GetUserMessageAfter(userID, sinceTime.Add(time.Nanosecond*100))
	if err != nil {
		log.Printf("get user message error: %v", err)
		return
	}

	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		manager: h.manager,
		conn:    conn,
		send:    utils.NewUnboundedChan[*entity.Message](2),
		userID:  userID,
		once:    sync.Once{},
	}
	h.manager.register <- client

	for _, msg := range pendingMessages {
		client.send.In <- &msg
	}

	go client.sendRoutine()
	go client.readRoutine()
}

func (h *WebSocketHost) clientManageRoutine() {
	log.Println("client manage routine start")
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
	return "/:user_id/host/conn"
}

func (h *WebSocketHost) HandlerMethod() string {
	return "GET"
}

func (h *WebSocketHost) Send(msg *types.Message) error {
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
	return "WebsocketHost"
}
