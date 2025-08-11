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
			log.Printf("Error reading message: %v", err)
			break
		}

		log.Printf("Received: %s", message)

		// Echo the message back to client
		err = conn.WriteMessage(messageType, message)
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
