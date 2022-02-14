package Sort

import (
	"Packages/src/pkg/MongoFilter/Sources"
)

type Settings struct {
	SortAscKey         string
	SortDescKey        string
	AcceptedSortFields map[string]string
	Sources            *Sources.Source
}

func NewSortSettings() *Settings {
	return &Settings{
		SortAscKey:         "asc",
		SortDescKey:        "dsc",
		AcceptedSortFields: map[string]string{},
		Sources:            Sources.New(Sources.Query),
	}
}
