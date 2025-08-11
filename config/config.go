package config

import (
	"flag"
	"fmt"
)

type Config struct {
	ServerMode string
	Hostname   string
	Port       string
	Version    string
}

const AppVersion = "0.0.1"

func NewConfig() *Config {
	// Default values
	defaultHostname := "192.168.1.0"
	defaultPort := "8000"
	defaultServerMode := "websocket"

	// Command-line flags
	hostname := flag.String("ip", defaultHostname, "IP address to bind the server")
	port := flag.String("port", defaultPort, "Port to bind the server")
	serverMode := flag.String("mode", defaultServerMode, "Server mode (websocket or http)")

	// Parse the arguments
	flag.Parse()

	return &Config{
		Hostname:   *hostname,
		Port:       *port,
		Version:    AppVersion,
		ServerMode: *serverMode,
	}
}

func (c *Config) GetFullServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Hostname, c.Port)
}
