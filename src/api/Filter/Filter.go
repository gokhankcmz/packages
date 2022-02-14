package Filter

import (
	"Packages/src/api/Type/ErrorTypes"
	"Packages/src/pkg/MongoFilter"
	"Packages/src/pkg/MongoFilter/Filters"
	"Packages/src/pkg/MongoFilter/Pagination"
	"Packages/src/pkg/MongoFilter/Sort"
	"Packages/src/pkg/MongoFilter/Sources"
	"time"
)

func GetFilter() *MongoFilter.MongoFilter {
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

	return mf
}
