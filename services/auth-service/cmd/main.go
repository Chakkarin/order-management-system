package main

import (
	"services/auth-service/internal/config"
	"services/auth-service/internal/server"
)

func main() {

	// Load Configuration
	cfg := config.LoadConfig()

	server := server.NewServer(cfg)

	server.Start()

}
