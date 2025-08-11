package websocket

import (
	"log"
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
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("Client connected: %s", conn.RemoteAddr())

	// Handle messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			// Check if it's a normal closure
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure) {
				log.Println("Client disconnected")
			} else {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		log.Printf("Received: %s. MessageType: %d", message, messageType)

		// Handle different message types
		switch messageType {
		case websocket.TextMessage:
			// Echo the text message
			err = conn.WriteMessage(websocket.TextMessage, []byte("Echo: "+string(message)))
		case websocket.BinaryMessage:
			// Echo the binary message
			err = conn.WriteMessage(websocket.BinaryMessage, message)
		case websocket.PingMessage:
			// Respond with Pong
			err = conn.WriteMessage(websocket.PongMessage, message)
		default:
			log.Printf("Unknown message type: %d", messageType)
			continue
		}

		if err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}

	log.Printf("Client disconnected: %s", conn.RemoteAddr())
}

// StartServer starts the WebSocket server
func StartServer(address string) error {
	http.HandleFunc("/ws", HandleWebSocket)
	log.Printf("WebSocket server starting on %s", address)
	return http.ListenAndServe(address, nil)
}
