package websocket

import (
	"sync"

	"notify-api/app/manager/push/item"
	"notify-api/db/ent/generate"
)

var manager = &clientManager{
	clients: make(map[int]map[*client]bool),
	lock:    sync.RWMutex{},
}

type clientManager struct {
	clients map[int]map[*client]bool
	lock    sync.RWMutex
}

func (m *clientManager) send(msg *item.PushMessage) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	clients, ok := m.clients[msg.User.ID]
	if ok {
		for c := range clients {
			c.send <- msg
		}
	}
}

func (m *clientManager) register(c *client) {
	m.lock.Lock()
	defer m.lock.Unlock()

	userID := c.context.User.ID
	clients, ok := m.clients[userID]
	if !ok {
		m.clients[userID] = make(map[*client]bool)
	}
	clients[c] = true
}

func (m *clientManager) unregister(c *client) {
	m.lock.Lock()
	defer m.lock.Unlock()

	userID := c.context.User.ID
	clients, ok := m.clients[userID]
	if ok {
		delete(clients, c)
		if len(clients) == 0 {
			delete(m.clients, userID)
		}
	}
}

func (m *clientManager) unregisterByDevice(d *generate.Device) {
	m.lock.Lock()
	defer m.lock.Unlock()

	userID := d.Edges.User.ID
	clients, ok := m.clients[userID]
	if ok {
		for c := range clients {
			if c.device.Identifier == d.Identifier {
				delete(clients, c)
			}
		}
		if len(clients) == 0 {
			delete(m.clients, userID)
		}
	}
}
