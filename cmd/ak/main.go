package main

import (
	"context"

	a "github.com/jobquestvault/platform-go-challenge/internal/app"
	c "github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	l "github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

const (
	appName = "asset-keeper"
)

func main() {
	cfg := c.Load()
	log := l.NewLogger(cfg.Log.Level)
	app := a.NewApp(appName, log, cfg)

	ctx := context.Background()
	app.Setup(ctx)
	err := app.Start(ctx)
	if err != nil {
		log.Error(err)
	}
}
