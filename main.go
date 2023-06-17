package main

import (
	"os"

	"github.com/laracarvalho/goboards/config"
	"github.com/laracarvalho/goboards/router"
)

var logger *config.Logger

func main() {
	logger := config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	r := router.Start()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)
}
