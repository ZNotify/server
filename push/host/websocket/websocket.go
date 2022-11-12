package websocket

import (
	"time"

	"notify-api/utils/user"

	pushTypes "notify-api/push/types"
)

type Host struct {
	manager *wsManager
}

func (h *Host) GetInitialTokenMeta() string {
	return time.Now().Format(time.RFC3339Nano)
}

func (h *Host) Start() error {
	go h.clientManageRoutine()
	return nil
}

func (h *Host) Init() error {
	h.manager = &wsManager{
		userClients: make(map[string]map[*wsClient]bool),
		register:    make(chan *wsClient),
		unregister:  make(chan *wsClient),
		broadcast:   make(chan *pushTypes.Message),
	}

	for _, v := range user.Users() {
		h.manager.userClients[v] = make(map[*wsClient]bool)
	}

	return nil
}

func (h *Host) Send(msg *pushTypes.Message) error {
	h.manager.broadcast <- msg
	return nil
}

func (h *Host) Name() string {
	return "WebSocketHost"
}
