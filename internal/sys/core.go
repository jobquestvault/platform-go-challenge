package sys

import (
	"context"

	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	Core interface {
		Log() log.Logger
		Cfg() *cfg.Config
		Lifecycle
	}

	BaseCore struct {
		log log.Logger
		cfg *cfg.Config
	}
)

func NewCore(log log.Logger, cfg *cfg.Config) *BaseCore {
	return &BaseCore{
		log: log,
		cfg: cfg,
	}
}

func (bs *BaseCore) Log() log.Logger {
	return bs.log
}

func (bs *BaseCore) Cfg() *cfg.Config {
	return bs.cfg
}

func (bc *BaseCore) Setup(ctx context.Context) error {
	bc.Log().Info("Default core setup")
	return nil
}

func (bc *BaseCore) Start(ctx context.Context) error {
	bc.Log().Info("Default core start")
	return nil
}

func (bc *BaseCore) Stop(ctx context.Context) error {
	bc.Log().Info("Default core stop")
	return nil
}
