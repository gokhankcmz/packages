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
		ThisSettingIsAMap: map[string]string{
			"f1": "v1",
			"v2": "v2",
		},
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
	},
}
