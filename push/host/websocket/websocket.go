package websocket

import (
	"strconv"

	"notify-api/utils/user"

	"notify-api/ent/dao"
	pushTypes "notify-api/push/item"
)

type Host struct {
	manager *wsManager
}

func (h *Host) GetDeviceInitialMeta() string {
	return strconv.FormatUint(uint64(dao.Message.GetLatestMessage().ID), 10)
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
		broadcast:   make(chan *pushTypes.PushMessage),
	}

	for _, v := range user.Users() {
		h.manager.userClients[v] = make(map[*wsClient]bool)
	}

	return nil
}

func (h *Host) Send(msg *pushTypes.PushMessage) error {
	h.manager.broadcast <- msg
	return nil
}

func (h *Host) Name() string {
	return "WebSocketHost"
}
