package MongoFilter

import (
	"Packages/src/pkg/MongoFilter/Filters"
	"Packages/src/pkg/MongoFilter/Pagination"
	"Packages/src/pkg/MongoFilter/Sort"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/url"
	"strconv"
)

type MongoFilter struct {
	PaginationSettings      *Pagination.Settings
	SortSettings            *Sort.Settings
	MultipleFilterSeparator string
	FilterItems             []Filters.IFilter
}

func NewBasicFilter() *MongoFilter {
	return &MongoFilter{
		PaginationSettings:      Pagination.NewPaginationSettings(),
		SortSettings:            Sort.NewSortSettings(),
		MultipleFilterSeparator: ",",
		FilterItems:             []Filters.IFilter{},
	}
}

func (mf *MongoFilter) AddFilters(filters ...Filters.IFilter) {
	for _, filter := range filters {
		mf.FilterItems = append(mf.FilterItems, filter)
	}
}

func (mf *MongoFilter) AddSort(Settings *Sort.Settings) {
	mf.SortSettings = Settings
}

func (mf *MongoFilter) AddPagination(Settings *Pagination.Settings) {
	mf.PaginationSettings = Settings
}

func (mf *MongoFilter) GetFilter(QueryParams url.Values, HeaderValues http.Header) *bson.M {
	filter := bson.M{}
	fields := mf.FilterItems
	for _, Item := range fields {
		QueryValue := Item.GetSourceValue(QueryParams, HeaderValues)
		if QueryValue != "" {
			Item.AddToBson(QueryValue, &filter)
		}
	}
	return &filter
}

func (mf *MongoFilter) GetFindOptions(QueryParams url.Values, HeaderValues http.Header) *options.FindOptions {
	findOptions := options.Find()
	showAll, err := strconv.ParseBool(mf.PaginationSettings.Sources.GetSourceValue(QueryParams, HeaderValues, mf.PaginationSettings.ShowAllKey))
	if !showAll || err != nil {

		ppValue, err := strconv.ParseInt(mf.PaginationSettings.Sources.GetSourceValue(QueryParams, HeaderValues, mf.PaginationSettings.PerPageKey), 10, 64)
		if err != nil || ppValue <= 0 {
			ppValue = mf.PaginationSettings.DefaultPerPage
		}
		if ppValue > mf.PaginationSettings.MaxPerPage {
			ppValue = mf.PaginationSettings.MaxPerPage
		}
		offsetValue, err := strconv.ParseInt(mf.PaginationSettings.Sources.GetSourceValue(QueryParams, HeaderValues, mf.PaginationSettings.OffsetKey), 10, 64)
		if err != nil || offsetValue < 0 {
			offsetValue = 0
		}
		findOptions.SetSkip(offsetValue)
		findOptions.SetLimit(ppValue)
	}

	if len(mf.SortSettings.AcceptedSortFields) != 0 {
		sort := bson.D{}
		for SourceFieldName, TargetFieldName := range mf.SortSettings.AcceptedSortFields {
			QueryValue := mf.SortSettings.Sources.GetSourceValue(QueryParams, HeaderValues, SourceFieldName)
			if QueryValue == mf.SortSettings.SortAscKey {
				sort = append(sort, bson.E{TargetFieldName, 1})
			}
			if QueryValue == mf.SortSettings.SortDescKey {
				sort = append(sort, bson.E{TargetFieldName, -1})
			}
		}
		findOptions.SetSort(sort)
	}

	return findOptions
}
