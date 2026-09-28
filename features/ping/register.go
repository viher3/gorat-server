package ping

import (
	"encoding/json"

	"github.com/viher3/gorat-server/network"
	"github.com/viher3/gorat-server/shared/commandbus"
	"github.com/viher3/gorat-server/shared/message"
)

// Register wires the ping use case into the command bus and the message router.
func Register(bus *commandbus.Bus, router *message.Router, clients *network.Clients) {
	h := NewHandler(clients)

	commandbus.RegisterHandler(bus, h.Handle)

	router.Register("ping", func(payload json.RawMessage) (commandbus.Command, error) {
		var cmd Command
		if err := json.Unmarshal(payload, &cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	})
}
