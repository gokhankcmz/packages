package Sort

import (
	"Packages/src/pkg/MongoFilter/Sources"
	"net/http"
	"net/url"
)

func (s *Settings) SetAscKey(ascKey string) *Settings {
	s.SortAscKey = ascKey
	return s
}

func (s *Settings) SetDescKey(descKey string) *Settings {
	s.SortDescKey = descKey
	return s
}

func (s *Settings) AddAcceptedSortField(FieldName, CorrespondingFieldInDB string) *Settings {
	s.AcceptedSortFields[FieldName] = CorrespondingFieldInDB
	return s
}

func (s *Settings) SetSource(sources ...func(QueryValues url.Values, headerValues http.Header, Keyword string) string) *Settings {
	s.Sources = Sources.New(sources...)
	return s
}
