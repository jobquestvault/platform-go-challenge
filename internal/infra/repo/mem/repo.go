package mem

import (
	"context"
	"math"

	"github.com/google/uuid"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
	"github.com/jobquestvault/platform-go-challenge/internal/infra/db"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
)

type favData struct {
	ID      string
	UserID  string
	AssetID string
}

type Favable interface {
	GetID() string
}

type InMemoryAssetRepo struct {
	Assets []model.Asset[model.Favable]
	Favs   []Favable
}

func NewInMemoryAssetRepo() *InMemoryAssetRepo {
	return &InMemoryAssetRepo{
		Assets: make([]model.Asset, 0),
		Favs:   make([]Favable, 0),
	}
}

func (ar *InMemoryAssetRepo) Tx(ctx context.Context) (db.Tx, error) {
	// Implement the transaction logic for in-memory storage if needed
	return nil, nil
}

func (ar *InMemoryAssetRepo) GetAssets(ctx context.Context, page, size int) ([]model.Asset, int, error) {
	start := (page - 1) * size
	end := start + size

	if start >= len(ar.Assets) {
		return nil, 0, nil
	}

	if end > len(ar.Assets) {
		end = len(ar.Assets)
	}

	assets := ar.Assets[start:end]
	totalPages := int(math.Ceil(float64(len(ar.Assets)) / float64(size)))

	return assets, totalPages, nil
}

func (ar *InMemoryAssetRepo) GetFaved(ctx context.Context, userID string, page, size int) ([]model.Asset, int, error) {
	favIDs := ar.getFavoriteIDsByUserID(userID)
	favAssets := make([]model.Asset, 0)

	for _, asset := range ar.Assets {
		if ar.isFav(asset.GetID(), favIDs) {
			favAssets = append(favAssets, asset)
		}
	}

	start := (page - 1) * size
	end := start + size

	if start >= len(favAssets) {
		return nil, 0, nil
	}

	if end > len(favAssets) {
		end = len(favAssets)
	}

	favAssets = favAssets[start:end]
	totalPages := int(math.Ceil(float64(len(favAssets)) / float64(size)))

	return favAssets, totalPages, nil
}

func (ar *InMemoryAssetRepo) AddFav(ctx context.Context, userID, assetID string) error {
	if ar.isFav(assetID, ar.getFavoriteIDsByUserID(userID)) {
		return errors.NewError("asset already favorited by the user")
	}

	favID := uuid.New().String()
	fav := &favData{
		ID:      favID,
		UserID:  userID,
		AssetID: assetID,
	}

	ar.Favs = append(ar.Favs, fav)
	return nil
}

func (ar *InMemoryAssetRepo) RemoveFav(ctx context.Context, userID, assetID string) error {
	favIDs := ar.getFavoriteIDsByUserID(userID)
	index := ar.findFavIndex(assetID, favIDs)

	if index == -1 {
		return errors.NewError("asset is not favorited by the user")
	}

	// Remove the favorite data from the slice
	ar.Favs = append(ar.Favs[:index], ar.Favs[index+1:]...)

	return nil
}

func (ar *InMemoryAssetRepo) findFavIndex(assetID string, favIDs []string) int {
	for i, favID := range favIDs {
		if favID == assetID {
			return i
		}
	}
	return -1
}

func (ar *InMemoryAssetRepo) isFav(assetID string, favIDs []string) bool {
	for _, favID := range favIDs {
		if favID == assetID {
			return true
		}
	}
	return false
}

func (ar *InMemoryAssetRepo) getFavoriteIDsByUserID(userID string) []string {
	favIDs := make([]string, 0)
	for _, fav := range ar.Favs {
		if fav.GetID() == userID {
			favIDs = append(favIDs, fav.AssetID)
		}
	}
	return favIDs
}
