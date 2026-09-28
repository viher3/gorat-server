package network

import clientmodel "github.com/viher3/gorat-server/network/model"

type Clients struct {
	clients map[int]*clientmodel.Client
}

func NewClients() *Clients {
	return &Clients{
		clients: make(map[int]*clientmodel.Client),
	}
}

func (c *Clients) AddClient(id int, client *clientmodel.Client) {
	c.clients[id] = client
}

func (c *Clients) RemoveClient(id int) {
	delete(c.clients, id)
}

func (c *Clients) GetClient(id int) (*clientmodel.Client, bool) {
	client, exists := c.clients[id]
	return client, exists
}

func (c *Clients) GetAllClients() map[int]*clientmodel.Client {
	return c.clients
}

func (c *Clients) Count() int {
	return len(c.clients)
}
