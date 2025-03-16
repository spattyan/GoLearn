package main

import (
	"GoLearn/config"
	"GoLearn/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v",err)
		return
	}

	// initialize router
	router.Initialize()
}
