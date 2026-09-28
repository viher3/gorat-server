// Package message decodes the JSON envelope received over the wire
// ({"type": "...", "payload": {...}}) into a concrete commandbus.Command.
package message

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/viher3/gorat-server/shared/commandbus"
)

type Envelope struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// Decoder turns a raw JSON payload into the concrete Command for one message type.
type Decoder func(payload json.RawMessage) (commandbus.Command, error)

type Router struct {
	mu       sync.RWMutex
	decoders map[string]Decoder
}

func NewRouter() *Router {
	return &Router{decoders: make(map[string]Decoder)}
}

// Register binds a wire type name to the Decoder that builds its Command.
func (r *Router) Register(msgType string, decode Decoder) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.decoders[msgType] = decode
}

// Decode parses the envelope and builds the Command for its type.
func (r *Router) Decode(raw []byte) (commandbus.Command, error) {
	var env Envelope
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, fmt.Errorf("message: invalid envelope: %w", err)
	}

	r.mu.RLock()
	decode, ok := r.decoders[env.Type]
	r.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("message: unknown type %q", env.Type)
	}
	return decode(env.Payload)
}
