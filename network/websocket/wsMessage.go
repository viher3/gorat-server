package websocket

import "encoding/json"

type WsMessage struct {
	ID      string                 `json:"id"`
	Payload map[string]interface{} `json:"payload"`
}

// NewWsMessage creates a new WebSocket message with the given ID and payload.
// The payload is a map that can hold any data structure.
func NewWsMessage(id string, payload map[string]interface{}) *WsMessage {
	if payload == nil {
		payload = make(map[string]interface{})
	}
	return &WsMessage{
		ID:      id,
		Payload: payload,
	}
}

// FromJSON parses a JSON string and returns a WsMessage struct.
// Returns the WsMessage and any unmarshaling error.
func NewWsMessageFromJSON(jsonStr string) (*WsMessage, error) {
	var msg WsMessage
	err := json.Unmarshal([]byte(jsonStr), &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
