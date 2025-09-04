package main

import (
	"os"
	"time"

	"github.com/widiskel/uomi-testnet-bot/internal/app"
	"github.com/widiskel/uomi-testnet-bot/internal/config"
	"github.com/widiskel/uomi-testnet-bot/internal/platform/logger"
	"github.com/widiskel/uomi-testnet-bot/internal/platform/ui"
)

func main() {
	_ = logger.Init("logs/app.log")
	defer logger.Close()

	ui.StartUISystem()
	defer ui.StartUISystem()

	cfg := config.Load()

	if err := app.New(cfg).Run(); err != nil {
		print(err.Error())
		os.Exit(1)
	}

	time.Sleep(1 * time.Second)
}
