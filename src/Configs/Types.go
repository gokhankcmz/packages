package Configs

import (
	"time"
)

type AppConfig struct {
	AppFirstRun             bool
	MongoConnectionURI      string
	DBName                  string
	CollectionName          string
	ApplicationName         string
	MongoConnectionDuration time.Duration
	LoggerSettings          LoggerSettings
	Filters                 []FilterObject
	Searchs                 []SearchObject
	PaginationSettings      PaginationSettings
	SortSettings            SortSettings
}

type LoggerSettings struct {
	PrintRequestInfo  bool
	PrintResponseInfo bool
	MaxRespBodySize   int64
	MaxRespDuration   int64
	LogToTerminal     bool
	LogLevelKeyword   string
	LogSuccessful     struct {
		Active   bool
		Loglevel string
	}
	LogClientErrors struct {
		Active   bool
		Loglevel string
	}
	LogServerErrors struct {
		Active   bool
		Loglevel string
	}
}

type FilterObject struct {
	SourceFieldName string
	Sources         string
	FieldType       string
	ComparisonOp    string
	TargetFieldName string
}

type SearchObject struct {
	SourceFieldName string
	Sources         string
	ComparisonOp    string
	TargetFieldName string
}

type PaginationSettings struct {
	Sources        string
	MaxPerPage     int64
	DefaultPerPage int64
	PerPageKey     string
	OffSetKey      string
	ShowAllKey     string
}

type SortSettings struct {
	Sources            string
	AscKey             string
	DescKey            string
	AcceptedSortFields []string
}
