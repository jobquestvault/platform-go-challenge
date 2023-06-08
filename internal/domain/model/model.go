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
	Data       []float64
	Favorite
}

type Insight struct {
	ID
	Text  string
	Topic string
	Favorite
}

type Audience struct {
	ID
	Gender              string
	BirthCountry        string
	AgeGroup            string
	HoursSpentOnSocial  int
	NumPurchasesLastMth int
	Favorite
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
		Data T
	}
)

func NewAsset[T Favable](id, name string, data T) Asset[Favable] {
	return Asset[Favable]{ID: ID{ID: id, Name: name}, Data: data}
}
