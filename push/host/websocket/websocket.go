package websocket

import (
	"notify-api/utils/user"

	"notify-api/db/entity"
	pushTypes "notify-api/push/types"
)

type Host struct {
	manager *wsManager
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
		broadcast:   make(chan *entity.Message),
	}

	for _, v := range user.Users() {
		h.manager.userClients[v] = make(map[*wsClient]bool)
	}

	return nil
}

func (h *Host) Send(msg *pushTypes.Message) error {
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

func (h *Host) Name() string {
	return "WebSocketHost"
}
