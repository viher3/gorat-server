package socket

import (
	"bufio"
	"errors"
	"io"
	"log/slog"
	"net"
)

func StartServer(address string, log *slog.Logger) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Info("socket server listening", "address", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("failed to accept connection", "error", err)
			continue
		}
		go handleConnection(conn, log)
	}
}

func handleConnection(conn net.Conn, log *slog.Logger) {
	defer conn.Close()

	connLog := log.With("remote_addr", conn.RemoteAddr().String())
	connLog.Info("client connected")
	defer connLog.Info("client disconnected")

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				connLog.Error("failed to read message", "error", err)
			}
			return
		}
		connLog.Info("message received", "message", message)
	}
}
