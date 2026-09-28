package ping

import (
	"context"
	"fmt"
	"strconv"

	"github.com/viher3/gorat-server/network"
	clientmodel "github.com/viher3/gorat-server/network/model"
)

type Handler struct {
	clients *network.Clients
}

func NewHandler(clients *network.Clients) *Handler {
	return &Handler{clients: clients}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) (any, error) {
	fmt.Println("Feature ping executed!")
	fmt.Println("param value " + cmd.ClientID)

	client, created := h.clients.GetOrCreate(cmd.ClientID, func() *clientmodel.Client {
		return clientmodel.NewClient(cmd.ClientID, cmd.ClientID, "TODO", "TODO")
	})
	if !created {
		client.Connected()
	}

	fmt.Println("total clients: " + strconv.Itoa(h.clients.Count()))
	fmt.Println("listing clients: ")
	for id, client := range h.clients.GetAllClients() {
		fmt.Printf("Client ID: %s, Client: %+v\n", id, client)
	}

	return nil, nil
}
