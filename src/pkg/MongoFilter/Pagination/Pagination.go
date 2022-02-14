package Pagination

import (
	"Packages/src/pkg/MongoFilter/Sources"
)

type Settings struct {
	DefaultPerPage int64
	MaxPerPage     int64
	OffsetKey      string
	PerPageKey     string
	ShowAllKey     string
	Sources        *Sources.Source
}

func NewPaginationSettings() *Settings {
	return &Settings{
		DefaultPerPage: 10,
		MaxPerPage:     50,
		OffsetKey:      "offset",
		PerPageKey:     "page",
		ShowAllKey:     "",
		Sources:        Sources.New(Sources.Query, Sources.Header),
	}
}
