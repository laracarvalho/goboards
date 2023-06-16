package main

import (
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

	router.Start()
}
