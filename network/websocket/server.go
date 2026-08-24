package websocket

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin
		return true
	},
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP connection to WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error("failed to upgrade connection", "error", err)
			return
		}
		defer conn.Close()

		connLog := log.With("remote_addr", conn.RemoteAddr().String())
		connLog.Info("client connected")

		// Handle messages
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				// Check if it's a normal closure
				if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure) {
					connLog.Info("client disconnected")
				} else {
					connLog.Error("failed to read message", "error", err)
				}
				break
			}

			connLog.Info("message received", "message", string(message), "message_type", messageType)

			// Handle different message types
			switch messageType {
			case websocket.TextMessage:
				wsMessage, wsMessageErr := NewWsMessageFromJSON(string(message))
				if wsMessageErr != nil {
					connLog.Error("failed to parse websocket message", "error", wsMessageErr)
					continue
				}
				connLog.Info("parsed websocket message", "id", wsMessage.ID, "payload", wsMessage.Payload)

				// Echo the text message
				err = conn.WriteMessage(websocket.TextMessage, []byte("Echo: "+string(message)))
			case websocket.BinaryMessage:
				// Echo the binary message
				err = conn.WriteMessage(websocket.BinaryMessage, message)
			case websocket.PingMessage:
				// Respond with Pong
				err = conn.WriteMessage(websocket.PongMessage, message)
			default:
				connLog.Warn("unknown message type", "message_type", messageType)
				continue
			}

			if err != nil {
				connLog.Error("failed to write message", "error", err)
				break
			}
		}

		connLog.Info("connection closed")
	}
}

// StartServer starts the WebSocket server
func StartServer(address string, log *slog.Logger) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", HandleWebSocket(log))
	log.Info("websocket server listening", "address", address)
	return http.ListenAndServe(address, mux)
}
