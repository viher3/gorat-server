package main

import (
	"fmt"

	"github.com/viher3/gorat-server/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println("GoRat Server v" + config.AppVersion)
	fmt.Println("Server is running at:", cfg.GetFullServerAddress())
}
