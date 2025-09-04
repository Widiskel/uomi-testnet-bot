package app

import (
	"sync"

	"github.com/widiskel/uomi-testnet-bot/internal/app/worker"
	"github.com/widiskel/uomi-testnet-bot/internal/config"
)

type App struct{ cfg config.Config }

func New(cfg config.Config) *App { return &App{cfg: cfg} }

func (app *App) Run() error {
	accounts, err := app.cfg.LoadAccounts()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for idx, acc := range accounts {
		wg.Add(1)
		go func(i int, a string) {
			defer wg.Done()
			worker.Run(a, i, app.cfg)
		}(idx, acc)
	}
	wg.Wait()
	return nil
}
