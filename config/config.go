package config

import "fmt"

type Config struct {
	Hostname string
	Port     string
	Version  string
}

const AppVersion = "0.0.1"

func NewConfig() *Config {
	return &Config{
		Hostname: "192.168.1.0",
		Port:     "8000",
		Version:  AppVersion,
	}
}

func (c *Config) GetFullServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Hostname, c.Port)
}
