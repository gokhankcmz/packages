package Sources

import (
	"net/http"
	"net/url"
)

type Source struct {
	Sources []func(QueryValues url.Values, headerValues http.Header, Keyword string) string
}

func New(sources ...func(QueryValues url.Values, headerValues http.Header, Keyword string) string) *Source {
	return &Source{Sources: sources}
}

func (s *Source) GetSourceValue(QueryValues url.Values, headerValues http.Header, Keyword string) string {
	for _, SourceFunc := range s.Sources {
		SourceValue := SourceFunc(QueryValues, headerValues, Keyword)
		if SourceValue != "" {
			return SourceValue
		}
	}
	return ""
}

func Header(QueryValues url.Values, headerValues http.Header, Keyword string) string {
	return headerValues.Get(Keyword)
}

func Query(QueryValues url.Values, headerValues http.Header, Keyword string) string {
	return QueryValues.Get(Keyword)

}
