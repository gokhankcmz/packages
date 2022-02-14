package Configs

import "time"

type AppConfig struct {
	AppFirstRun             bool
	MongoConnectionURI      string
	DBName                  string
	CollectionName          string
	ApplicationName         string
	MongoConnectionDuration time.Duration
}
