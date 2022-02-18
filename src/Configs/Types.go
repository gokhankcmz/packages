package Configs

import "time"

type AppConfig struct {
	AppFirstRun             bool
	MongoConnectionURI      string
	DBName                  string
	CollectionName          string
	ApplicationName         string
	MongoConnectionDuration time.Duration
	LoggerSettings          LoggerSettings
	ThisSettingIsAMap       map[string]string
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
