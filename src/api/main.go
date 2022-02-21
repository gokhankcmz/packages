package main

import (
	"Packages/src/Configs"
	"Packages/src/api/Filter"
	"Packages/src/api/Handlers"
	"Packages/src/api/HealthChecks"
	"Packages/src/api/Logging"
	"Packages/src/api/Mongo"
	UserRepository "Packages/src/api/Repository"
	consul "Packages/src/pkg/Consul"
	KvpConverter "Packages/src/pkg/KVP"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Application Starting...")
	e := echo.New()

	Configs.SetDefault("dev")
	config := Configs.GetConfigs()

	Logging.Init(config)
	logger := Logging.GetLogger()

	//Mongo
	mongoClient, err := Mongo.GetMongoClient(config)
	if err != nil {
		logger.CreateLog().LogFatal(0)
	}
	mongoCollection, err := Mongo.GetMongoDbCollection(mongoClient, config)
	if err != nil {
		logger.CreateLog().LogFatal(0)
	}

	repo := UserRepository.NewRepository(mongoCollection)
	health := HealthChecks.GetHealthChecks(mongoClient, config)
	filter := Filter.SetHardcodedFilter()
	consulClient := consul.Init("consul:8500")
	kvpConverter := KvpConverter.New("/", config.ApplicationName)
	//First run settings
	if config.AppFirstRun {
		repo.CreateInitialData(30)
		configKVP := kvpConverter.GetKVP(config)
		for k, v := range configKVP {
			_ = consulClient.Set(k, v)
		}
	}

	Handlers.NewHandler(e, repo, filter, health, &consulClient, kvpConverter)
	e.Start(":80")
}
