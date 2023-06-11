package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (c Chart) MarshalJSON() ([]byte, error) {
	last := len(c.Data) - 1

	data := "{"
	for i, val := range c.Data {
		n := strconv.Itoa(int(val))
		if i < last {
			data = fmt.Sprintf("%s%s, ", data, n)
		} else {
			data = fmt.Sprintf("%s%s", data, n)
		}
	}
	data = data + "}"

	m := struct {
		ID         string
		Type       string
		Title      string
		XAxisTitle string
		YAxisTitle string
		Data       string
	}{
		ID:         c.GetID(),
		Type:       c.GetType(),
		Title:      c.Title,
		XAxisTitle: c.XAxisTitle,
		YAxisTitle: c.YAxisTitle,
		Data:       data,
	}

	return json.Marshal(m)
}

func (a Asset[T]) MarshalJSON() ([]byte, error) {
	m := struct {
		ID          string
		Type        string
		UserID      string
		AssetID     string
		Name        string
		Description string
		Data        T
	}{
		ID:          a.ID.ID,
		Type:        a.Data.GetType(),
		UserID:      a.UserID,
		AssetID:     a.Data.GetID(),
		Name:        a.ID.Name,
		Description: a.Description,
		Data:        a.Data,
	}

	return json.Marshal(m)
}
