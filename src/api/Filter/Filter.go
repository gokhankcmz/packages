package Filter

import (
	"Packages/src/Configs"
	"Packages/src/api/Type/ErrorTypes"
	"Packages/src/pkg/MongoFilter"
	"Packages/src/pkg/MongoFilter/Filters"
	"Packages/src/pkg/MongoFilter/Pagination"
	"Packages/src/pkg/MongoFilter/Sort"
	"Packages/src/pkg/MongoFilter/Sources"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var filter *MongoFilter.MongoFilter

func GetFilter() *MongoFilter.MongoFilter {
	return filter
}

func SetHardcodedFilter() *MongoFilter.MongoFilter {
	mf := MongoFilter.NewBasicFilter()
	mf.AddFilters(
		Filters.NewFilter().
			SetSourceFieldName("minupdatedat").
			SetSources(Sources.Query, Sources.Header).
			SetAsTimeField(time.RFC3339).
			MustBeGreaterThanOrEqualTo().
			SetTargetFieldInDB("document.updatedat").
			SetPanicAndSave(ErrorTypes.InvalidModel),
		Filters.NewRegexSearch().
			SetSourceFieldName("name").
			SetSources(Sources.Query).
			SetComparisonOption("i").
			SetTargetFieldInDB("name").
			SetPanicAndSave(ErrorTypes.InvalidModel),
		Filters.NewRegexSearch().
			SetSourceFieldName("email").
			SetSources(Sources.Query, Sources.Header).
			SetComparisonOption("i").
			SetTargetFieldInDB("email").
			SetPanicAndSave(ErrorTypes.InvalidModel),
	)

	mf.AddPagination(
		Pagination.NewPaginationSettings().
			SetSources(Sources.Query, Sources.Header).
			SetMaxPerPage(2).
			SetDefaultPerPage(1).
			SetPerPageKey("perpage").
			SetOffSetKey("page"),
	)

	mf.AddSort(
		Sort.NewSortSettings().
			SetSource(Sources.Query, Sources.Header).
			SetAscKey("asc").SetDescKey("dsc").
			AddAcceptedSortField("sortage", "age").
			AddAcceptedSortField("sortcreationtime", "document.createdat"),
	)
	filter = mf
	return mf
}

func SetDynamicFilter(config *Configs.AppConfig) *MongoFilter.MongoFilter {
	mf := MongoFilter.NewBasicFilter()
	for _, v := range config.Searchs {
		mf.AddFilters(
			Filters.NewRegexSearch().SetSourceFieldName(v.SourceFieldName).
				SetSources(getSources(v.Sources)...).
				SetComparisonOption(v.ComparisonOp).
				SetTargetFieldInDB(v.TargetFieldName).
				SetPanicAndSave(ErrorTypes.InvalidModel))
	}

	for _, v := range config.Filters {
		mf.AddFilters(
			Filters.NewFilter().SetSourceFieldName(v.SourceFieldName).
				SetSources(getSources(v.Sources)...).
				SetFieldTypeCorrectionFunc(getFieldTypeFunc(v.FieldType)).
				SetComparisonOption(v.ComparisonOp).
				SetTargetFieldInDB(v.TargetFieldName).
				SetPanicAndSave(ErrorTypes.InvalidModel))
	}
	mf.AddPagination(
		Pagination.NewPaginationSettings().
			SetSources(getSources(config.PaginationSettings.Sources)...).
			SetMaxPerPage(config.PaginationSettings.MaxPerPage).
			SetOffSetKey(config.PaginationSettings.OffSetKey).
			SetDefaultPerPage(config.PaginationSettings.DefaultPerPage).
			SetPerPageKey(config.PaginationSettings.PerPageKey).
			SetShowAllKey(config.PaginationSettings.ShowAllKey),
	)
	mf.AddSort(
		Sort.NewSortSettings().
			SetSource(getSources(config.SortSettings.Sources)...).
			SetAscKey(config.SortSettings.AscKey).
			SetDescKey(config.SortSettings.DescKey),
	)
	for _, v := range config.SortSettings.AcceptedSortFields {
		sf := strings.Split(strings.TrimSpace(v), ",")
		mf.SortSettings.AddAcceptedSortField(sf[0], sf[1])
	}
	filter = mf
	return mf
}

func getSources(sources string) []func(QueryValues url.Values, headerValues http.Header, Keyword string) string {
	sourcesStringArray := strings.Split(strings.TrimSpace(strings.ToLower(sources)), ",")
	var sourceFuncs []func(QueryValues url.Values, headerValues http.Header, Keyword string) string
	for _, v := range sourcesStringArray {
		switch v {
		case "query":
			sourceFuncs = append(sourceFuncs, Sources.Query)
		case "header":
			sourceFuncs = append(sourceFuncs, Sources.Header)
		}
	}
	return sourceFuncs
}
func getFieldTypeFunc(fieldType string) func(stringValue string) interface{} {
	ft := strings.ToLower(fieldType)
	switch ft {
	case "integer", "int":
		return Filters.Integer
	case "time", "time.rfc3339", "rfc3339", "339":
		return Filters.DateTime(time.RFC3339)
	}
	return nil
}
