package Configs

import "time"

var DefaultConfigs = map[string]AppConfig{
	"dev": {
		AppFirstRun:             true,
		MongoConnectionURI:      "mongodb://mongo:27017",
		DBName:                  "PackageApiDB",
		ApplicationName:         "PackagesApi",
		CollectionName:          "Users",
		MongoConnectionDuration: time.Second * 5,
		LoggerSettings: LoggerSettings{
			PrintRequestInfo:  true,
			PrintResponseInfo: true,
			MaxRespBodySize:   1000,
			MaxRespDuration:   1000,
			LogToTerminal:     true,
			LogLevelKeyword:   "level",
			LogSuccessful: struct {
				Active   bool
				Loglevel string
			}{true, "info"},
			LogClientErrors: struct {
				Active   bool
				Loglevel string
			}{true, "info"},
			LogServerErrors: struct {
				Active   bool
				Loglevel string
			}{true, "error"},
		},
		Filters: []FilterObject{
			{
				SourceFieldName: "minupdatedat",
				Sources:         "Query, Header",
				FieldType:       "time.RFC3339",
				ComparisonOp:    "$gte",
				TargetFieldName: "document.updatedat",
			},
		},
		Searchs: []SearchObject{
			{
				SourceFieldName: "name",
				Sources:         "Query, Header",
				ComparisonOp:    "i",
				TargetFieldName: "name",
			},
			{
				SourceFieldName: "email",
				Sources:         "Query, Header",
				ComparisonOp:    "i",
				TargetFieldName: "email",
			},
		},
		PaginationSettings: PaginationSettings{
			Sources:        "Query, Header",
			MaxPerPage:     10,
			DefaultPerPage: 5,
			PerPageKey:     "perpage",
			OffSetKey:      "offset",
		},
		SortSettings: SortSettings{
			Sources:            "query,header",
			AscKey:             "asc",
			DescKey:            "desc",
			AcceptedSortFields: []string{"sortage,age", "sortcreationtime,document.createdat"},
		},
	},
}
