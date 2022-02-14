package Filters

import (
	"Packages/src/pkg/MongoFilter/Sources"
	"net/http"
	"net/url"
)

func (f *Filter) setSourceFieldName(SourceFieldName string) {
	f.SourceFieldName = SourceFieldName
}
func (f *Filter) setTargetFieldName(TargetFieldName string) {
	f.TargetFieldName = TargetFieldName
}
func (f *Filter) setCompOp(ComparisonOperator string) {
	f.ComparisonOperator = ComparisonOperator
}
func (f *Filter) setSources(sources ...func(QueryValues url.Values, headerValues http.Header, Keyword string) string) {
	f.Sources = Sources.New(sources...)
}
func (f *Filter) setFieldCorrFunc(fnc func(StringValue string) interface{}) {
	f.FieldTypeCorrectionFunction = fnc
}
func (f *Filter) setPanic(err error) {
	f.PanicIfFails = err
}

func NewFilter() *Filter {
	return &Filter{
		AbstractFilter:              &AbstractFilter{},
		FieldTypeCorrectionFunction: nil,
	}
}

type Step1 struct{ filter *Filter }

func (f *Filter) SetSourceFieldName(SourceFieldName string) *Step1 {
	f.setSourceFieldName(SourceFieldName)
	return &Step1{
		filter: f,
	}
}

type Step2 struct{ filter *Filter }

func (s1 *Step1) SetSources(sources ...func(QueryValues url.Values, headerValues http.Header, Keyword string) string) *Step2 {
	s1.filter.setSources(sources...)
	return &Step2{filter: s1.filter}
}

type Step3 struct{ filter *Filter }

func (s2 *Step2) SetAsIntegerField() *Step3 {
	s2.filter.setFieldCorrFunc(Integer)
	return &Step3{filter: s2.filter}
}

func (s2 *Step2) SetAsTimeField(TimeLayout string) *Step3 {
	s2.filter.setFieldCorrFunc(DateTime(TimeLayout))
	return &Step3{filter: s2.filter}
}

func (s2 *Step2) SetFieldTypeCorrectionFunc(f func(stringValue string) interface{}) *Step3 {
	s2.filter.setFieldCorrFunc(f)
	return &Step3{filter: s2.filter}
}

type Step4 struct{ filter *Filter }

func (s3 *Step3) SetComparisonOption(op string) *Step4 {
	s3.filter.setCompOp(op)
	return &Step4{filter: s3.filter}
}

func (s3 *Step3) MustBeGreaterThanOrEqualTo() *Step4 {
	s3.filter.setCompOp(MustBeGreaterThanOrEqualTo)
	return &Step4{filter: s3.filter}
}

func (s3 *Step3) MustBeLessThanOrEqualTo() *Step4 {
	s3.filter.setCompOp(MustBeLessThanOrEqualTo)
	return &Step4{filter: s3.filter}
}

type Step5 struct{ filter *Filter }

func (s4 *Step4) SetTargetFieldInDB(CorrespondingFieldInDB string) *Step5 {
	s4.filter.setTargetFieldName(CorrespondingFieldInDB)
	return &Step5{filter: s4.filter}
}

func (s5 *Step5) SetPanicAndSave(PanicIfFails error) Filter {
	s5.filter.setPanic(PanicIfFails)
	return *s5.filter
}
