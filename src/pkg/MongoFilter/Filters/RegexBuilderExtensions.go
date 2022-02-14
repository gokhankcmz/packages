package Filters

import (
	"Packages/src/pkg/MongoFilter/Sources"
	"net/http"
	"net/url"
)

func NewRegexSearch() *RegexFilter {
	return &RegexFilter{}
}
func (f *RegexFilter) setSourceFieldName(SourceFieldName string) {
	f.SourceFieldName = SourceFieldName
}
func (f *RegexFilter) setTargetFieldName(TargetFieldName string) {
	f.TargetFieldName = TargetFieldName
}
func (f *RegexFilter) setCompOp(ComparisonOperator string) {
	f.ComparisonOperator = ComparisonOperator
}
func (f *RegexFilter) setSources(sources ...func(QueryValues url.Values, headerValues http.Header, Keyword string) string) {
	f.Sources = Sources.New(sources...)
}
func (f *RegexFilter) setPanic(err error) {
	f.PanicIfFails = err
}

type RegexStep1 struct{ filter *RegexFilter }

func (f *RegexFilter) SetSourceFieldName(SourceFieldName string) *RegexStep1 {
	f.SourceFieldName = SourceFieldName
	return &RegexStep1{filter: f}
}

type RegexStep2 struct{ filter *RegexFilter }

func (s1 *RegexStep1) SetSources(sources ...func(QueryValues url.Values, headerValues http.Header, Keyword string) string) *RegexStep2 {
	s1.filter.setSources(sources...)
	return &RegexStep2{filter: s1.filter}
}

type RegexStep3 struct{ filter *RegexFilter }

func (s2 *RegexStep2) SetComparisonOption(op string) *RegexStep3 {
	s2.filter.setCompOp(op)
	return &RegexStep3{filter: s2.filter}
}

type RegexStep4 struct{ filter *RegexFilter }

func (s3 *RegexStep3) SetTargetFieldInDB(CorrespondingFieldInDB string) *RegexStep4 {
	s3.filter.setTargetFieldName(CorrespondingFieldInDB)
	return &RegexStep4{filter: s3.filter}
}

func (s4 *RegexStep4) SetPanicAndSave(PanicIfFails error) *RegexFilter {
	s4.filter.setPanic(PanicIfFails)
	return s4.filter
}
