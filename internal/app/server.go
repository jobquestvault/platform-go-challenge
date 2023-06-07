package app

import (
	"encoding/json"
	"net/http"
)

var assets = []Asset[Favorite]{
	NewAsset("e54d8fde-ec90-494b-8f40-e66acfd40fec",
		"sample-chart",
		Chart{
			ID:         "23fefba1-7276-4f09-8bb7-ea3decea1700",
			Title:      "Revenue",
			XAxisTitle: "Time",
			YAxisTitle: "Revenue",
			Data:       []float64{1, 20, 2, 21, 3, 22, 4, 23, 5, 24, 6, 25, 7, 26, 8, 27, 9, 28, 10, 29},
			favorite:   true,
		},
	),
	NewAsset("c7a5d9c8-7d0c-456d-89a5-399b77e5cc79",
		"sample-insight",
		Insight{
			ID:       "56723ff2-a222-4927-9c6d-c0875b208b9e",
			Text:     "Lumos Nexus Solutions",
			Topic:    "Digital Transformation Strategies",
			favorite: true,
		},
	),
	NewAsset("c89394c8-b035-4184-bd88-ae1de08a7e31",
		"sample-audience",
		Audience{
			ID:                  "e284fd63-784d-431b-9372-786b6f3a21f6",
			Gender:              "female",
			BirthCountry:        "uk",
			AgeGroup:            "20-30",
			HoursSpentOnSocial:  2,
			NumPurchasesLastMth: 1,
			favorite:            true,
		},
	),
}

func (s *Server) favorites(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(assets)
	if err != nil {
		// TODO: Implement after defining handling strategy
	}
}

func (s *Server) addFavorite(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/add" endpoint
	panic("not implemented yet")
}

func (s *Server) removeFavorite(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/remove" endpoint
	panic("not implemented yet")
}

func (s *Server) updateFavorite(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handling for "/favorites/edit" endpoint
	panic("not implemented yet")
}
