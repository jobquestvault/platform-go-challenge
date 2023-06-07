package model

type (
	Favorite interface {
		Faved() bool
	}

	favorite bool
)

type Chart struct {
	ID         string
	Title      string
	XAxisTitle string
	YAxisTitle string
	Data       []float64
	favorite
}

func (i favorite) Faved() bool {
	return bool(i)
}

type Insight struct {
	ID    string
	Text  string
	Topic string
	favorite
}

type Audience struct {
	ID                  string
	Gender              string
	BirthCountry        string
	AgeGroup            string
	HoursSpentOnSocial  int
	NumPurchasesLastMth int
	favorite
}

type Asset[T Favorite] struct {
	ID   string
	Name string
	Data T
}

func NewAsset[T Favorite](id, name string, data T) Asset[Favorite] {
	return Asset[Favorite]{ID: id, Name: name, Data: data}
}
