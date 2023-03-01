package websocket

import (
	"strconv"
	"sync"
	"time"

	"notify-api/app/common"
	"notify-api/app/manager/push/item"
	"notify-api/app/models"
	"notify-api/db/dao"
	"notify-api/db/ent/generate"

	"github.com/fasthttp/websocket"
	"go.uber.org/zap"
)

type client struct {
	conn    *websocket.Conn
	send    chan *item.PushMessage
	once    sync.Once
	context *common.Context
	onClose func(*client)
	device  *generate.Device
}

func newClient(ctx *common.Context, device *generate.Device, conn *websocket.Conn, onClose func(*client)) *client {
	return &client{
		context: ctx,
		conn:    conn,
		send:    make(chan *item.PushMessage, 1),
		onClose: onClose,
		device:  device,
	}
}

func (c *client) run() {
	go c.writeRoutine()
	go c.readRoutine()
}

func (c *client) updateReadDeadline() {
	_ = c.conn.SetReadDeadline(time.Now().Add(pongTimeout))
}

func (c *client) updateWriteDeadline() {
	_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
}

func (c *client) pong() error {
	err := c.conn.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(writeWait))
	if err != nil {
		zap.S().Errorf("client %d %s write control error: %v", c.context.User.ID, c.device.Identifier, err)
	}
	return err
}

func (c *client) ping() error {
	err := c.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait))
	if err != nil {
		zap.S().Errorf("client %d %s write control error: %v", c.context.User.ID, c.device.Identifier, err)
	}
	return err
}

func (c *client) writeMessage(msg *item.PushMessage) error {
	c.updateWriteDeadline()
	err := c.conn.WriteJSON(models.FromPushMessage(*msg))
	if err != nil {
		c.logWebsocketError(err)
	}
	return err
}

func (c *client) pongHandler(string) error {
	c.updateReadDeadline()
	return nil
}

func (c *client) pingHandler(string) error {
	return c.pong()
}

func (c *client) logWebsocketError(err error) {
	normalCodes := []int{websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseInternalServerErr}
	if websocket.IsUnexpectedCloseError(err, normalCodes...) {
		zap.S().Infof("client %d %s read routine exit %v", c.context.User.ID, c.device.Identifier, err)
	}
}

func (c *client) writeRoutine() {
	ticket := time.NewTicker(pingPeriod)
	defer func() {
		ticket.Stop()
		c.close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				return
			}

			err := c.writeMessage(msg)
			if err != nil {
				zap.S().Warnw("client write message error", "error", err)
				return
			}

			dao.Device.UpdateDeviceChannelMeta(c.context, c.device, strconv.Itoa(msg.SequenceID))

		case <-ticket.C:
			err := c.ping()
			if err != nil {
				zap.S().Warnw("client ping error", "error", err)
				return
			}

		}
	}
}

func (c *client) readRoutine() {
	defer func() {
		c.close()
	}()

	c.conn.SetReadLimit(maxMessageSize)

	c.updateReadDeadline()

	c.conn.SetPongHandler(c.pongHandler)
	c.conn.SetPingHandler(c.pingHandler)

	for {
		_, _, err := c.conn.NextReader()
		if err != nil {
			c.logWebsocketError(err)
			return
		}
	}
}

func (c *client) close() {
	zap.S().Debugf("client %d %s try close", c.context.User.ID, c.device.Identifier)
	c.once.Do(func() {
		zap.S().Debugf("client %d %s real close", c.context.User.ID, c.device.Identifier)
		_ = c.conn.Close()
		close(c.send)
		c.onClose(c)
	})
}
