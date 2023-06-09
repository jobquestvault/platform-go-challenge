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
		GetAssets(ctx context.Context, userId string, status ...AssetStatus) (assets []model.Asset[model.Favable], err error)
		AddFav(ctx context.Context, userID, assetType, ID string) error
		RemoveFav(ctx context.Context, userID, assetType, ID string) error
		UpdateFav(ctx context.Context, userID, assetType, ID, name string) error
	}

	AssetStatus string
)

const (
	Faved    AssetStatus = "Faved"
	NotFaved AssetStatus = "NotFaved"
)
