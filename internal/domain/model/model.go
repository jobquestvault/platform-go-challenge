package model

import (
	"reflect"
	"strings"
)

type (
	Identifiable interface {
		GetID() string
	}

	Nameable interface {
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

type (
	Chart struct {
		ID
		Title      string
		XAxisTitle string
		YAxisTitle string
		Data       []uint8
	}

	Insight struct {
		ID
		Text  string
		Topic string
	}

	Audience struct {
		ID
		Gender                string
		BirthCountry          string
		AgeGroup              string
		HoursSpentOnSocial    int
		NumPurchasesLastMonth int
	}
)

type (
	Typable interface {
		GetType() string
	}
)

// GetType returns the type of the struct.
// Reflection can be replaced by a hardcoded string.
func (c Chart) GetType() string {
	typeObj := reflect.TypeOf(c)
	return strings.ToLower(typeObj.Name())
}

func (i Insight) GetType() string {
	typeObj := reflect.TypeOf(i)
	return strings.ToLower(typeObj.Name())
}

func (a Audience) GetType() string {
	typeObj := reflect.TypeOf(a)
	return strings.ToLower(typeObj.Name())
}

type (
	Favable interface {
		Identifiable
		Typable
	}

	Faved interface {
		Favable
		IsFaved() bool
	}
)

type Asset[T Favable] struct {
	ID
	Type        string
	UserID      string
	AssetID     string
	Name        string
	Description string
	Data        T
}

func NewAsset[T Favable](id, name, description, assetType string, data T) Asset[Favable] {
	return Asset[Favable]{ID: ID{ID: id, Name: name}, Type: assetType, Description: description, Data: data}
}

func (a *Asset[T]) GetType() string {
	return a.Data.GetType()
}

func (a *Asset[T]) IsFaved() bool {
	return a.UserID != "" &&
		a.AssetID != "" &&
		a.Type != ""
}
