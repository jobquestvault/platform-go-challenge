package service

import (
	"github.com/jobquestvault/platform-go-challenge/internal/domain/port"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	AssetService interface {
		Repo() port.Repo
	}

	Asset struct {
		sys.Core
		repo port.Repo
	}
)

func NewService(repo port.Repo, log log.Logger, cfg *cfg.Config) *Asset {
	return &Asset{
		Core: sys.NewCore(log, cfg),
		repo: repo,
	}
}

func (a Asset) Repo() port.Repo {
	return a.repo
}
