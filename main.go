package main

import (
	"os"

	"github.com/viher3/gorat-server/config"
	"github.com/viher3/gorat-server/features/ping"
	"github.com/viher3/gorat-server/network"
	"github.com/viher3/gorat-server/network/socket"
	"github.com/viher3/gorat-server/shared/commandbus"
	"github.com/viher3/gorat-server/shared/logger"
	"github.com/viher3/gorat-server/shared/message"
)

func main() {
	appLog := logger.New("app")
	netLog := logger.New("network")

	cfg := config.NewConfig()
	appLog.Info("starting gorat-server", "version", config.AppVersion, "address", cfg.GetFullServerAddress(), "mode", cfg.ServerMode)

	bus := commandbus.New()
	router := message.NewRouter()
	clients := network.NewClients()

	ping.Register(bus, router, clients)

	var err error
	switch cfg.ServerMode {
	case "socket":
		err = socket.StartServer(cfg.GetFullServerAddress(), netLog, bus, router)
	default:
		appLog.Error("unknown server mode", "mode", cfg.ServerMode)
		os.Exit(1)
	}

	appLog.Error("server stopped", "mode", cfg.ServerMode, "error", err)
	os.Exit(1)
}
