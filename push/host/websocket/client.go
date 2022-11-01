package websocket

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"notify-api/db/entity"
	"notify-api/db/model"
	"notify-api/utils/ds"
)

const (
	writeWait = 10 * time.Second

	timeout = 30 * time.Second

	pingPeriod = (timeout * 7) / 10

	maxMessageSize = 512
)

type wsClient struct {
	manager *wsManager

	conn *websocket.Conn

	send *ds.UnboundedChan[*entity.Message]

	userID   string
	deviceID string

	once sync.Once
}

type wsManager struct {
	userClients map[string]map[*wsClient]bool

	register chan *wsClient

	unregister chan *wsClient

	broadcast chan *entity.Message
}

type wsMessage entity.Message

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *wsClient) writeRoutine() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.close()
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
				c.close()
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
				c.close()
				return
			}
		}
	}
}

func (c *wsClient) readRoutine() {
	defer func() {
		c.close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	for {
		_, _, err := c.conn.NextReader()
		if err != nil {
			normalCodes := []int{websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseInternalServerErr}
			if websocket.IsUnexpectedCloseError(err, normalCodes...) {
				zap.S().Infof("wsClient %s %s read routine exit %v", c.userID, c.deviceID[0:7], err)
			}
			break
		}
	}
}

func (c *wsClient) close() {
	zap.S().Debugf("wsClient %s %s try close", c.userID, c.deviceID[0:7])
	c.once.Do(func() {
		zap.S().Debugf("wsClient %s %s real close", c.userID, c.deviceID[0:7])
		_ = c.conn.Close()
		c.manager.unregister <- c
		close(c.send.In)
	})
}

func (h *Host) clientManageRoutine() {
	zap.S().Debug("wsClient manage routine start")
	deleteClient := func(client *wsClient) {
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
