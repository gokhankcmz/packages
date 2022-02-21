package Configs

import (
	"fmt"
	"strconv"
	"time"
)

var configurations AppConfig

func GetConfigs() *AppConfig {
	return &configurations
}

func SetDefault(env string) {
	configurations = DefaultConfigs[env]
}
func Set(config AppConfig) {
	configurations = config
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

	//FilterSettings
	for i := 0; ; i++ {
		pfx := fmt.Sprintf("%v/Filters/%v", prefix, i)
		if _, ok := kvp[pfx]; ok {
			configurations.Filters = append(configurations.Filters, FilterObject{
				SourceFieldName: kvp[pfx+"SourceFieldName"],
				Sources:         kvp[pfx+"Sources"],
				FieldType:       kvp[pfx+"SourceFieldName"],
				ComparisonOp:    kvp[pfx+"ComparisonOp"],
				TargetFieldName: kvp[pfx+"TargerFieldName"],
			})
		} else {
			break
		}
	}

	//SearchSettings
	for i := 0; ; i++ {
		pfx := fmt.Sprintf("%v/Searchs/%v", prefix, i)
		if _, ok := kvp[pfx]; ok {
			configurations.Searchs = append(configurations.Searchs, SearchObject{
				SourceFieldName: kvp[pfx+"SourceFieldName"],
				Sources:         kvp[pfx+"Sources"],
				ComparisonOp:    kvp[pfx+"ComparisonOp"],
				TargetFieldName: kvp[pfx+"TargetFieldName"],
			})
		} else {
			break
		}
	}

	//PaginationSettings
	dpp, _ := strconv.ParseInt(kvp[prefix+"/PaginationSettings/DefaultPerPage"], 10, 64)
	configurations.PaginationSettings.DefaultPerPage = dpp
	mpp, _ := strconv.ParseInt(kvp[prefix+"/PaginationSettings/MaxPerPage"], 10, 64)
	configurations.PaginationSettings.MaxPerPage = mpp
	configurations.PaginationSettings.OffSetKey = kvp[prefix+"/PaginationSettings/OffSetKey"]
	configurations.PaginationSettings.PerPageKey = kvp[prefix+"/PaginationSettings/PerPageKey"]
	configurations.PaginationSettings.ShowAllKey = kvp[prefix+"/PaginationSettings/ShowAllKey"]
	configurations.PaginationSettings.Sources = kvp[prefix+"/PaginationSettings/Sources"]

	//SortSettings

	configurations.SortSettings.Sources = kvp[prefix+"/SortSettings/Sources"]
	configurations.SortSettings.AscKey = kvp[prefix+"/SortSettings/AscKey"]
	configurations.SortSettings.DescKey = kvp[prefix+"/SortSettings/DescKey"]
	for i := 0; ; i++ {
		pfx := fmt.Sprintf("%v/SortSettings/AcceptedSortFields/%v", prefix, i)
		if v, ok := kvp[pfx]; ok {
			configurations.SortSettings.AcceptedSortFields = append(configurations.SortSettings.AcceptedSortFields, v)
		} else {
			break
		}
	}
	return &configurations

}
