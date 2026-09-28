package ping

import (
	"context"

	"github.com/viher3/gorat-server/network"
)

type Handler struct {
	clients *network.Clients
}

func NewHandler(clients *network.Clients) *Handler {
	return &Handler{clients: clients}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) (any, error) {
	// TODO: caso de uso real (p.ej. localizar el cliente por ClientID y
	// marcarlo como conectado con client.Connected()).
	return nil, nil
}
