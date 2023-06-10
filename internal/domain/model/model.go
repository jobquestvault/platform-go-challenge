package model

type (
	Identifiable interface {
		GetID() string
		GetName() string
	}

	ID struct {
		ID   string
		Name string
	}
)

func (id ID) GetID() string {
	return id.ID
}

func (id ID) GetName() string {
	return id.Name
}

type Chart struct {
	ID
	Title      string
	XAxisTitle string
	YAxisTitle string
	Data       []uint8
}

type Insight struct {
	ID
	Text  string
	Topic string
}

type Audience struct {
	ID
	Gender              string
	BirthCountry        string
	AgeGroup            string
	HoursSpentOnSocial  int
	NumPurchasesLastMth int
}

type Asset struct {
	ID
	UserID      string
	AssetID     string
	AssetType   string
	Name        string
	Description string
}

type (
	Favable interface {
		Faved() bool
	}

	Favorite bool
)

func (i Favorite) Faved() bool {
	return bool(i)
}

type (
	FavableAsset interface {
		Identifiable
		Favable
	}

	Asset[T Favable] struct {
		ID
		Type string
		Data T
	}
)

func NewAsset[T Favable](id, name, assetType string, data T) Asset[Favable] {
	return Asset[Favable]{ID: ID{ID: id, Name: name}, Type: assetType, Data: data}
}
