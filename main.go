package main

import (
	"fmt"
	"log"

	"github.com/viher3/gorat-server/config"
	"github.com/viher3/gorat-server/network/websocket"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println("GoRat Server v" + config.AppVersion)
	fmt.Println("Server is running at:", cfg.GetFullServerAddress())

	// Start the WebSocket server
	if cfg.ServerMode == "websocket" {
		log.Fatal(websocket.StartServer(cfg.GetFullServerAddress()))
	}

}
