package socket

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net"
	"strings"

	"github.com/viher3/gorat-server/shared/commandbus"
	"github.com/viher3/gorat-server/shared/message"
)

func StartServer(address string, log *slog.Logger, bus *commandbus.Bus, router *message.Router) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Info("socket server listening", "address", address)

	// init clients manager
	//clients := connectedClients.NewClients()
	//log.Info("clients manager initialized", "count", clients.Count())

	// accept connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("failed to accept connection", "error", err)
			continue
		}
		go handleConnection(conn, log, bus, router)
	}
}

func handleConnection(conn net.Conn, log *slog.Logger, bus *commandbus.Bus, router *message.Router) {
	defer conn.Close()

	connLog := log.With("remote_addr", conn.RemoteAddr().String())
	connLog.Info("client connected")
	defer connLog.Info("client disconnected")

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				connLog.Error("failed to read message", "error", err)
			}
			return
		}
		parsedMessage := strings.TrimSuffix(line, "\n")
		var envelope message.Envelope
		if err := json.Unmarshal([]byte(parsedMessage), &envelope); err != nil {
			connLog.Error("failed to parse message JSON", "error", err)
			continue
		}

		cmd, err := router.Decode([]byte(parsedMessage))

		if err != nil {
			connLog.Error("failed to decode message", "error", err)
			continue
		}

		if _, err := bus.Dispatch(context.Background(), cmd); err != nil {
			connLog.Error("failed to dispatch command", "error", err)
		}
	}
}
