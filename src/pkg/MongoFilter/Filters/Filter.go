package Filters

import (
	"Packages/src/pkg/MongoFilter/Sources"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/url"
)

type IFilter interface {
	AddToBson(httpValue string, filter *bson.M) *bson.M
	GetSourceValue(QueryValues url.Values, headerValues http.Header) string
}

type AbstractFilter struct {
	SourceFieldName    string
	TargetFieldName    string
	ComparisonOperator string
	Sources            *Sources.Source
	PanicIfFails       error
}

type Filter struct {
	*AbstractFilter
	FieldTypeCorrectionFunction func(StringValue string) interface{}
}

func (f Filter) AddToBson(httpValue string, filter *bson.M) *bson.M {
	(*filter)[f.TargetFieldName] = bson.M{
		f.ComparisonOperator: f.FieldTypeCorrectionFunction(httpValue),
	}
	return filter
}
func (f Filter) GetSourceValue(QueryValues url.Values, headerValues http.Header) string {
	return f.Sources.GetSourceValue(QueryValues, headerValues, f.SourceFieldName)
}

type RegexFilter AbstractFilter

func (f RegexFilter) AddToBson(httpValue string, filter *bson.M) *bson.M {
	(*filter)[f.TargetFieldName] = bson.M{
		"$regex": primitive.Regex{
			Pattern: httpValue,
			Options: f.ComparisonOperator,
		},
	}
	return filter
}
func (f RegexFilter) GetSourceValue(QueryValues url.Values, headerValues http.Header) string {
	return f.Sources.GetSourceValue(QueryValues, headerValues, f.SourceFieldName)
}
