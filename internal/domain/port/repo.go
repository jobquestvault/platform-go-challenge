package port

import (
	"context"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
	"github.com/jobquestvault/platform-go-challenge/internal/infra/db"
)

type (
	Repo interface {
		Tx(ctx context.Context) (tx db.Tx, err error)
	}

	AssetRepo interface {
		Repo
		GetAssets(ctx context.Context) ([]model.Asset[model.Favorite], error)
		GetFaved(ctx context.Context) ([]model.Asset[model.Favorite], error)
		AddFav(ctx context.Context, asset *model.Asset[model.Favorite]) error
		RemoveFav(ctx context.Context, asset *model.Asset[model.Favorite]) error
		UpdateFav(ctx context.Context, asset *model.Asset[model.Favorite]) error
	}
)
