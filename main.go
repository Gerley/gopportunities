package main

import (
	"github.com/Gerley/gopportunities/config"
	"github.com/Gerley/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLooger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}

// https://www.youtube.com/watch?v=wyEYpX5U4Vg
