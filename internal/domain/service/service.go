package service

import (
	"github.com/jobquestvault/platform-go-challenge/internal/domain/port"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	AssetService interface {
		Repo() port.AssetRepo
	}

	Asset struct {
		sys.Core
		repo port.AssetRepo
	}
)

func NewService(repo port.AssetRepo, log log.Logger, cfg *cfg.Config) *Asset {
	return &Asset{
		Core: sys.NewCore(log, cfg),
		repo: repo,
	}
}

func (a Asset) Repo() port.AssetRepo {
	return a.repo
}
