package network

import "time"

type Client struct {
	ID          string
	Name        string
	PublicIP    string
	PrivateIP   string
	isConnected bool
	lastSeen    int64
}

func NewClient(id, name, publicIP, privateIP string) *Client {
	return &Client{
		ID:          id,
		Name:        name,
		PublicIP:    publicIP,
		PrivateIP:   privateIP,
		isConnected: true,
		lastSeen:    time.Now().Unix(),
	}
}

func (c *Client) CheckIsConnected() bool {
	return c.lastSeen != 0 && time.Since(time.Unix(c.lastSeen, 0)) < 10*time.Minute
}

func (c *Client) Disconnected() {
	c.isConnected = false
}

func (c *Client) Connected() {
	c.isConnected = true
	c.lastSeen = time.Now().Unix()
}
