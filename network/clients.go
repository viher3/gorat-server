package network

import (
	"sync"

	clientmodel "github.com/viher3/gorat-server/network/model"
)

type Clients struct {
	mu      sync.RWMutex
	clients map[string]*clientmodel.Client
}

func NewClients() *Clients {
	return &Clients{
		clients: make(map[string]*clientmodel.Client),
	}
}

func (c *Clients) AddClient(id string, client *clientmodel.Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.clients[id] = client
}

func (c *Clients) RemoveClient(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.clients, id)
}

func (c *Clients) GetClient(id string) (*clientmodel.Client, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	client, exists := c.clients[id]
	return client, exists
}

// GetOrCreate returns the existing client for id, or atomically creates one
// with create and stores it if none exists yet. created is true when a new
// client was just added.
func (c *Clients) GetOrCreate(id string, create func() *clientmodel.Client) (client *clientmodel.Client, created bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if existing, exists := c.clients[id]; exists {
		return existing, false
	}
	client = create()
	c.clients[id] = client
	return client, true
}

// GetAllClients returns a snapshot copy of the current clients, safe to
// range over without holding the internal lock.
func (c *Clients) GetAllClients() map[string]*clientmodel.Client {
	c.mu.RLock()
	defer c.mu.RUnlock()
	snapshot := make(map[string]*clientmodel.Client, len(c.clients))
	for id, client := range c.clients {
		snapshot[id] = client
	}
	return snapshot
}

func (c *Clients) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.clients)
}
