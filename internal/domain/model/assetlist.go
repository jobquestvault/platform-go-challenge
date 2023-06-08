package model

// Assets is a simple list of assets to be used in initial phase of development.
// Will be removed later.
var Assets = []Asset[Favable]{
	NewAsset("e54d8fde-ec90-494b-8f40-e66acfd40fec",
		"sample-chart",
		Chart{
			ID:         ID{ID: "23fefba1-7276-4f09-8bb7-ea3decea1700", Name: "Sample Chart"},
			Title:      "Revenue",
			XAxisTitle: "Time",
			YAxisTitle: "Revenue",
			Data:       []float64{1, 20, 2, 21, 3, 22, 4, 23, 5, 24, 6, 25, 7, 26, 8, 27, 9, 28, 10, 29},
			Favorite:   true,
		},
	),
	NewAsset("c7a5d9c8-7d0c-456d-89a5-399b77e5cc79",
		"sample-insight",
		Insight{
			ID:       ID{ID: "56723ff2-a222-4927-9c6d-c0875b208b9e", Name: "Sample Insight"},
			Text:     "Lumos Nexus Solutions",
			Topic:    "Digital Transformation Strategies",
			Favorite: true,
		},
	),
	NewAsset("c89394c8-b035-4184-bd88-ae1de08a7e31",
		"sample-audience",
		Audience{
			ID:                  ID{ID: "e284fd63-784d-431b-9372-786b6f3a21f6", Name: "Sample Audience"},
			Gender:              "female",
			BirthCountry:        "uk",
			AgeGroup:            "20-30",
			HoursSpentOnSocial:  2,
			NumPurchasesLastMth: 1,
			Favorite:            true,
		},
	),
}
