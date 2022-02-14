package Configs

import "time"

var DefaultConfigs = map[string]AppConfig{
	"dev": {
		AppFirstRun:             true,
		MongoConnectionURI:      "mongodb://localhost:27017",
		DBName:                  "PackageApiDB",
		ApplicationName:         "PackageApi",
		CollectionName:          "Users",
		MongoConnectionDuration: time.Second * 5,
	},
}
