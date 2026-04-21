package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	clients map[string][]*websocket.Conn
	mu      sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[string][]*websocket.Conn),
	}
}

func (m *Manager) AddClient(channel string, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[channel] = append(m.clients[channel], conn)
}

func (m *Manager) RemoveClient(channel string, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	conns := m.clients[channel]
	for i, c := range conns {
		if c == conn {
			m.clients[channel] = append(conns[:i], conns[i+1:]...)
			break
		}
	}

	if len(m.clients[channel]) == 0 {
		delete(m.clients, channel)
	}
}

func (m *Manager) Send(channel string, data any) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, conn := range m.clients[channel] {
		_ = conn.WriteJSON(data)
	}
}

func (m *Manager) CloseChannel(channel string) {
	m.mu.Lock()
	conns, ok := m.clients[channel]
	if !ok {
		m.mu.Unlock()
		return
	}
	delete(m.clients, channel)
	m.mu.Unlock()

	for _, conn := range conns {
		_ = conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "done"),
		)
		_ = conn.Close()
	}
}
