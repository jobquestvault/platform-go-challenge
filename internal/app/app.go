package app

import (
	"context"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/service"
	"github.com/jobquestvault/platform-go-challenge/internal/infra/db/pg"
	"github.com/jobquestvault/platform-go-challenge/internal/infra/http"
	pgr "github.com/jobquestvault/platform-go-challenge/internal/infra/repo/pg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type App struct {
	sys.Core
	name    string
	http    *http.Server
	service service.AssetService
}

func NewApp(name string, log *log.BaseLogger, cfg *cfg.Config) *App {
	return &App{
		Core: sys.NewCore(log, cfg),
		name: name,
	}
}

func (app *App) Setup(ctx context.Context) error {
	log := app.Log()
	cfg := app.Cfg()

	// Databases
	db := pg.NewDB(log, cfg)
	err := db.Connect()
	if err != nil {
		return errors.Wrap("app setup error", err)
	}

	// Repo
	repo := pgr.NewAssetRepo(db, log, cfg)

	// Services
	svc := service.NewService(repo, log, cfg)

	// HTTP Server
	app.http = http.NewServer(svc, log, cfg)

	return app.http.Setup(ctx)
}

func (app *App) Start(ctx context.Context) error {
	return app.http.Start(ctx)
}

func (app *App) SetService(as service.AssetService) {
	app.service = as
}
