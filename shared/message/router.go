// Package message decodes the JSON envelope received over the wire
// ({"action": "...", "payload": {...}}) into a concrete commandbus.Command.
package message

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/viher3/gorat-server/shared/commandbus"
)

type Envelope struct {
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
}

// Decoder turns a raw JSON payload into the concrete Command for one message action.
type Decoder func(payload json.RawMessage) (commandbus.Command, error)

type Router struct {
	mu       sync.RWMutex
	decoders map[string]Decoder
}

func NewRouter() *Router {
	return &Router{decoders: make(map[string]Decoder)}
}

// Register binds a wire action name to the Decoder that builds its Command.
func (r *Router) Register(msgAction string, decode Decoder) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.decoders[msgAction] = decode
}

// Decode parses the envelope and builds the Command for its action.
func (r *Router) Decode(raw []byte) (commandbus.Command, error) {
	var env Envelope
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, fmt.Errorf("message: invalid envelope: %w", err)
	}

	r.mu.RLock()
	decode, ok := r.decoders[env.Action]
	r.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("message: unknown action %q", env.Action)
	}
	return decode(env.Payload)
}
