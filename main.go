package main

import (
	"os"

	"github.com/viher3/gorat-server/config"
	"github.com/viher3/gorat-server/network/socket"
	"github.com/viher3/gorat-server/network/websocket"
	"github.com/viher3/gorat-server/shared/logger"
)

func main() {
	appLog := logger.New("app")
	netLog := logger.New("network")

	cfg := config.NewConfig()
	appLog.Info("starting gorat-server", "version", config.AppVersion, "address", cfg.GetFullServerAddress(), "mode", cfg.ServerMode)

	var err error
	switch cfg.ServerMode {
	case "websocket":
		err = websocket.StartServer(cfg.GetFullServerAddress(), netLog)
	case "socket":
		err = socket.StartServer(cfg.GetFullServerAddress(), netLog)
	default:
		appLog.Error("unknown server mode", "mode", cfg.ServerMode)
		os.Exit(1)
	}

	appLog.Error("server stopped", "mode", cfg.ServerMode, "error", err)
	os.Exit(1)
}
