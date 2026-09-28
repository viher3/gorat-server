// Package commandbus routes Commands to the Handler registered for their
// concrete type, decoupling use cases from whatever transport triggers them.
package commandbus

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

// Command is a marker interface implemented by every use case input.
type Command interface{}

// Handler executes a single Command and returns an optional result.
type Handler func(ctx context.Context, cmd Command) (any, error)

type Bus struct {
	mu       sync.RWMutex
	handlers map[reflect.Type]Handler
}

func New() *Bus {
	return &Bus{handlers: make(map[reflect.Type]Handler)}
}

// Register binds every Command of cmd's concrete type to handler.
func (b *Bus) Register(cmd Command, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[reflect.TypeOf(cmd)] = handler
}

// Dispatch routes cmd to its registered handler.
func (b *Bus) Dispatch(ctx context.Context, cmd Command) (any, error) {
	b.mu.RLock()
	handler, ok := b.handlers[reflect.TypeOf(cmd)]
	b.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("commandbus: no handler registered for %T", cmd)
	}
	return handler(ctx, cmd)
}

// RegisterHandler is a type-safe helper so feature code never deals with `any`.
func RegisterHandler[C Command](b *Bus, handler func(ctx context.Context, cmd C) (any, error)) {
	var zero C
	b.Register(zero, func(ctx context.Context, cmd Command) (any, error) {
		return handler(ctx, cmd.(C))
	})
}
