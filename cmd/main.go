package main

import (
	"github.com/valdir-alves3000/go-opportunities/config"
	"github.com/valdir-alves3000/go-opportunities/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.SetupRouter()
}
