package app

import (
	"context"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/service"
	"github.com/jobquestvault/platform-go-challenge/internal/infra/http"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type App struct {
	sys.Core
	name    string
	server  *http.Server
	Service service.AssetService
}

func NewApp(name string, log *log.BaseLogger, cfg *cfg.Config) *App {
	return &App{
		Core:   sys.NewCore(log, cfg),
		name:   name,
		server: http.NewServer(log, cfg),
	}
}

func (app *App) Setup(ctx context.Context) {
	app.server.Setup(ctx)
}

func (app *App) Start(ctx context.Context) error {
	return app.server.Start(ctx)
}
