package Pagination

import (
	"Packages/src/pkg/MongoFilter/Sources"
	"net/http"
	"net/url"
)

func (pg *Settings) SetMaxPerPage(maxpp int64) *Settings {
	pg.MaxPerPage = maxpp
	return pg
}

func (pg *Settings) SetDefaultPerPage(defaultpp int64) *Settings {
	pg.DefaultPerPage = defaultpp
	return pg
}

func (pg *Settings) SetShowAllKey(key string) *Settings {
	pg.ShowAllKey = key
	return pg
}

func (pg *Settings) SetPerPageKey(key string) *Settings {
	pg.PerPageKey = key
	return pg
}

func (pg *Settings) SetOffSetKey(key string) *Settings {
	pg.OffsetKey = key
	return pg
}

func (pg *Settings) SetSources(sources ...func(QueryValues url.Values, headerValues http.Header, Keyword string) string) *Settings {
	pg.Sources = Sources.New(sources...)
	return pg
}
