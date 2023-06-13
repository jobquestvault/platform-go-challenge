package pg_test

import (
	"context"
	"testing"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
	pgdb "github.com/jobquestvault/platform-go-challenge/internal/infra/db/pg"
	"github.com/jobquestvault/platform-go-challenge/internal/infra/repo/pg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

func TestGetAssets(t *testing.T) {
	testCases := []struct {
		description    string
		userID         string
		expectedOutput []model.Asset[model.Favable]
	}{
		{
			description:    "User has no favorites",
			userID:         "user1",
			expectedOutput: []model.Asset[model.Favable]{},
		},
		{
			description: "User has favorite charts",
			userID:      "user2",
			expectedOutput: []model.Asset[model.Favable]{
				{
					ID:          model.ID{ID: "chart1", Name: "Chart 1"},
					Type:        "chart",
					UserID:      "user2",
					AssetID:     "asset1",
					Name:        "Favorite Chart 1",
					Description: "This is a favorite chart",
					Data: model.Chart{
						ID:         model.ID{ID: "chart1", Name: "Chart 1"},
						Title:      "Sales",
						XAxisTitle: "Time",
						YAxisTitle: "Revenue",
						Data:       []uint8{1, 2, 3, 4, 5},
					},
				},
			},
		},
		{
			description: "User has favorite insights",
			userID:      "user3",
			expectedOutput: []model.Asset[model.Favable]{
				{
					ID:          model.ID{ID: "insight1", Name: "Insight 1"},
					Type:        "insight",
					UserID:      "user3",
					AssetID:     "asset2",
					Name:        "Favorite Insight 1",
					Description: "This is a favorite insight",
					Data: model.Insight{
						ID:    model.ID{ID: "insight1", Name: "Insight 1"},
						Text:  "Lorem ipsum dolor sit amet",
						Topic: "Analysis",
					},
				},
				{
					ID:          model.ID{ID: "insight2", Name: "Insight 2"},
					Type:        "insight",
					UserID:      "user3",
					AssetID:     "asset3",
					Name:        "Favorite Insight 2",
					Description: "This is another favorite insight",
					Data: model.Insight{
						ID:    model.ID{ID: "insight2", Name: "Insight 2"},
						Text:  "Consectetur adipiscing elit",
						Topic: "Trends",
					},
				},
			},
		},
		{
			description: "User has favorite audiences",
			userID:      "user4",
			expectedOutput: []model.Asset[model.Favable]{
				{
					ID:          model.ID{ID: "audience1", Name: "Audience 1"},
					Type:        "audience",
					UserID:      "user4",
					AssetID:     "asset4",
					Name:        "Favorite Audience 1",
					Description: "This is a favorite audience",
					Data: model.Audience{
						ID:                    model.ID{ID: "audience1", Name: "Audience 1"},
						Gender:                "Male",
						BirthCountry:          "USA",
						AgeGroup:              "18-24",
						HoursSpentOnSocial:    2,
						NumPurchasesLastMonth: 5,
					},
				},
			},
		},
		{
			description: "User has favorites of different types",
			userID:      "user5",
			expectedOutput: []model.Asset[model.Favable]{
				{
					ID:          model.ID{ID: "chart2", Name: "Chart 2"},
					Type:        "chart",
					UserID:      "user5",
					AssetID:     "asset5",
					Name:        "Favorite Chart 2",
					Description: "This is another favorite chart",
					Data: model.Chart{
						ID:         model.ID{ID: "chart2", Name: "Chart 2"},
						Title:      "Expenses",
						XAxisTitle: "Category",
						YAxisTitle: "Amount",
						Data:       []uint8{10, 20, 30, 40, 50},
					},
				},
				{
					ID:          model.ID{ID: "insight3", Name: "Insight 3"},
					Type:        "insight",
					UserID:      "user5",
					AssetID:     "asset6",
					Name:        "Favorite Insight 3",
					Description: "This is another favorite insight",
					Data: model.Insight{
						ID:    model.ID{ID: "insight3", Name: "Insight 3"},
						Text:  "Ut enim ad minim veniam",
						Topic: "Strategy",
					},
				},
			},
		},
	}

	repo := pg.NewAssetRepo(getDb(), getLog(), getCfg())

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assets, _, err := repo.GetAssets(context.TODO(), 1, 14)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// Compare the expected output with the actual output
			if len(assets) != len(tc.expectedOutput) {
				t.Errorf("unexpected number of assets, got %d, want %d", len(assets), len(tc.expectedOutput))
			} else {
				for i := range assets {
					if assets[i] != tc.expectedOutput[i] {
						t.Errorf("unexpected asset at index %d, got %v, want %v", i, assets[i], tc.expectedOutput[i])
					}
				}
			}
		})
	}
}

func getDb() *pgdb.DB {
	logger := getLog()
	config := getCfg()

	postgresDB := pgdb.NewDB(logger, config)

	return postgresDB
}
func getLog() log.Logger {
	return log.NewLogger("debug")
}

func getCfg() *cfg.Config {
	config := &cfg.Config{
		Log: &cfg.LogConfig{
			Level: "info",
		},
		Server: &cfg.ServerConfig{
			Host: "localhost",
			Port: 8080,
		},
		DB: &cfg.DBConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "admin",
			Password: "password",
			Schema:   "ak",
			Name:     "pgc",
			SSL:      true,
		},
		Prop: &cfg.PropConfig{
			PageSize: 12,
		},
	}

	return config
}
