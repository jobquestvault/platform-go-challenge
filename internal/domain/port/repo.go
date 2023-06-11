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
		GetAssets(ctx context.Context, page, size int) (assets []model.Asset[model.Favable], pagesCount int, err error)
		GetFaved(ctx context.Context, userID string, page, size int) (assets []model.Asset[model.Favable], pagesCount int, err error)
		AddFav(ctx context.Context, userID, assetType, ID string) error
		UpdateFav(ctx context.Context, userID, assetType, ID, name, description string) error
		RemoveFav(ctx context.Context, userID, assetType, ID string) error
	}
)
