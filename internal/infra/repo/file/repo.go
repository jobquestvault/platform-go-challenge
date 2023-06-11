package file

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
	"github.com/jobquestvault/platform-go-challenge/internal/infra/db"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
)

type FileAssetRepo struct {
	assetFile string
	favFile   string
}

type FavData struct {
	ID      string
	UserID  string
	AssetID string
}

func NewFileAssetRepo() (*FileAssetRepo, error) {
	assetFile, err := generateTempFileName()
	if err != nil {
		return nil, err
	}

	favFile, err := generateTempFileName()
	if err != nil {
		return nil, err
	}

	return &FileAssetRepo{
		assetFile: assetFile,
		favFile:   favFile,
	}, nil
}

func generateTempFileName() (string, error) {
	file, err := os.CreateTemp("", "assets-")
	if err != nil {
		return "", errors.NewError("failed to generate temp file")
	}

	return file.Name(), nil
}

func (far *FileAssetRepo) Tx(ctx context.Context) (db.Tx, error) {
	return nil, nil
}

func (fr *FileAssetRepo) GetAssets(ctx context.Context, page, size int) ([]model.Asset[model.Favable], int, error) {
	start := (page - 1) * size
	end := start + size

	// Read assets from the asset file
	assets, err := fr.readAssetsFromFile()
	if err != nil {
		return nil, 0, err
	}

	// Perform pagination
	if start >= len(assets) {
		return nil, 0, nil
	}
	if end > len(assets) {
		end = len(assets)
	}
	assets = assets[start:end]

	// Calculate total pages
	totalPages := (len(assets) + size - 1) / size

	return assets, totalPages, nil
}

func (fr *FileAssetRepo) GetFaved(ctx context.Context, userID string, page, size int) ([]model.Asset[model.Favable], int, error) {
	start := (page - 1) * size
	end := start + size

	favs, err := fr.readFavoritesFromFile()
	if err != nil {
		return nil, 0, err
	}

	favIDs := make([]string, 0)
	for _, fav := range favs {
		if fav.UserID == userID {
			favIDs = append(favIDs, fav.AssetID)
		}
	}

	assets, err := fr.readAssetsFromFile()
	if err != nil {
		return nil, 0, err
	}

	favAssets := make([]model.Asset[model.Favable], 0)
	for _, asset := range assets {
		if contains(favIDs, asset.GetID()) {
			favAssets = append(favAssets, asset)
		}
	}

	if start >= len(favAssets) {
		return nil, 0, nil
	}
	if end > len(favAssets) {
		end = len(favAssets)
	}
	favAssets = favAssets[start:end]

	totalPages := (len(favAssets) + size - 1) / size

	return favAssets, totalPages, nil
}

func (fr *FileAssetRepo) AddFav(ctx context.Context, userID, assetType, assetID string) error {
	favs, err := fr.readFavoritesFromFile()
	if err != nil {
		return err
	}

	if isFavorited(favs, userID, assetID) {
		return errors.NewError("asset already favorited by the user")
	}

	fav := FavData{
		UserID:  userID,
		AssetID: assetID,
	}

	favs = append(favs, fav)

	err = fr.writeFavoritesToFile(favs)
	if err != nil {
		return err
	}

	return nil
}

func (far *FileAssetRepo) UpdateFav(ctx context.Context, userID, assetType, ID, name, description string) error {
	assets, err := far.loadAssets()
	if err != nil {
		return err
	}

	var found bool
	for i, asset := range assets {
		if asset.ID == ID && asset.Type == assetType {
			// Update the asset's name and description
			assets[i].Name = name
			assets[i].Description = description
			found = true
			break
		}
	}

	if !found {
		return errors.NewError("asset not found")
	}

	err = far.saveAssets(assets)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FileAssetRepo) RemoveFav(ctx context.Context, userID, assetType, assetID string) error {
	favs, err := fr.readFavoritesFromFile()
	if err != nil {
		return err
	}

	index := -1
	for i, fav := range favs {
		if fav.UserID == userID && fav.AssetID == assetID {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.NewError("asset is not favorited by the user")
	}

	favs = append(favs[:index], favs[index+1:]...)

	err = fr.writeFavoritesToFile(favs)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FileAssetRepo) readAssetsFromFile() ([]model.Asset[model.Favable], error) {
	assetBytes, err := ioutil.ReadFile(fr.assetFile)
	if err != nil {
		return nil, err
	}

	var assets []model.Asset[model.Favable]
	err = json.Unmarshal(assetBytes, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (fr *FileAssetRepo) readFavoritesFromFile() ([]FavData, error) {
	favBytes, err := ioutil.ReadFile(fr.favFile)
	if err != nil {
		return nil, err
	}

	var favorites []FavData
	err = json.Unmarshal(favBytes, &favorites)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func (fr *FileAssetRepo) writeFavoritesToFile(favorites []FavData) error {
	favBytes, err := json.Marshal(favorites)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fr.favFile, favBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func isFavorited(favs []FavData, userID, assetID string) bool {
	for _, fav := range favs {
		if fav.UserID == userID && fav.AssetID == assetID {
			return true
		}
	}
	return false
}

func contains(ids []string, id string) bool {
	for _, item := range ids {
		if item == id {
			return true
		}
	}
	return false
}

func (fr *FileAssetRepo) loadAssets(filePath string) ([]model.Asset[model.Favable], error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var assets []model.Asset
	err = json.Unmarshal(data, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (fr *FileAssetRepo) saveAssets(assets []model.Asset[model.Favable], filePath string) error {
	data, err := json.Marshal(assets)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
