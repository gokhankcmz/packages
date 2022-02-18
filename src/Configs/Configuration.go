package Configs

import (
	"strconv"
	"strings"
	"time"
)

var configurations AppConfig

func GetConfigs() *AppConfig {
	return &configurations
}

func SetDefault(env string) {
	configurations = DefaultConfigs[env]
}

func SetFromConsul(prefix string, kvp map[string]string) *AppConfig {
	configurations.ApplicationName = kvp[prefix+"/ApplicationName"]
	configurations.DBName = kvp[prefix+"/DBName"]
	configurations.MongoConnectionURI = kvp[prefix+"/MongoConnectionURI"]
	configurations.CollectionName = kvp[prefix+"/CollectionName"]
	fr, _ := strconv.ParseBool(kvp[prefix+"/AppFirstRun"])
	configurations.AppFirstRun = fr
	mcd, _ := time.ParseDuration(kvp[prefix+"/MongoConnectionDuration"])
	configurations.MongoConnectionDuration = mcd

	configurations.LoggerSettings.LogLevelKeyword = kvp[prefix+"/LoggerSettings/LogLevelKeyword"]
	lce, _ := strconv.ParseBool(kvp[prefix+"/LoggerSettings/LogClientErrors/Active"])
	configurations.LoggerSettings.LogClientErrors.Active = lce
	configurations.LoggerSettings.LogClientErrors.Loglevel = kvp[prefix+"/LoggerSettings/LogClientErrors/Loglevel"]
	lse, _ := strconv.ParseBool(kvp[prefix+"/LoggerSettings/LogServerErrors/Active"])
	configurations.LoggerSettings.LogServerErrors.Active = lse
	configurations.LoggerSettings.LogServerErrors.Loglevel = kvp[prefix+"/LoggerSettings/LogServerErrors/Loglevel"]
	lss, _ := strconv.ParseBool(kvp[prefix+"/LoggerSettings/LogSuccessful/Active"])
	configurations.LoggerSettings.LogSuccessful.Active = lss
	configurations.LoggerSettings.LogSuccessful.Loglevel = kvp[prefix+"/LoggerSettings/LogSuccessful/Loglevel"]
	ltt, _ := strconv.ParseBool(kvp[prefix+"/LoggerSettings/LogToTerminal"])
	configurations.LoggerSettings.LogToTerminal = ltt
	mbs, _ := strconv.ParseInt(kvp[prefix+"/LoggerSettings/MaxRespBodySize"], 10, 64)
	configurations.LoggerSettings.MaxRespBodySize = mbs
	mrd, _ := strconv.ParseInt(kvp[prefix+"/LoggerSettings/MaxRespDuration"], 10, 64)
	configurations.LoggerSettings.MaxRespDuration = mrd
	prqi, _ := strconv.ParseBool(kvp[prefix+"/LoggerSettings/PrintRequestInfo"])
	configurations.LoggerSettings.PrintRequestInfo = prqi
	prpi, _ := strconv.ParseBool(kvp[prefix+"/LoggerSettings/PrintResponseInfo"])
	configurations.LoggerSettings.PrintResponseInfo = prpi
	m := map[string]string{}
	for k, v := range kvp {
		if strings.Contains(k, prefix+"/ThisSettingIsAMap") {
			m[k] = v
		}
	}
	configurations.ThisSettingIsAMap = m
	return &configurations

}

type Test struct {
	Field1 string
	Field2 bool
	Field3 int64
}

var Testicin Test

func GetTest() *Test {
	return &Testicin
}

func SetTest() {
	Testicin = Test{
		Field1: "f1",
		Field2: true,
		Field3: 5,
	}
}
