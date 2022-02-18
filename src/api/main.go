package main

import (
	"Packages/src/Configs"
	"Packages/src/api/Filter"
	"Packages/src/api/Handlers"
	"Packages/src/api/HealthChecks"
	"Packages/src/api/Logging"
	"Packages/src/api/Middlewares"
	"Packages/src/api/Mongo"
	UserRepository "Packages/src/api/Repository"
	consul "Packages/src/pkg/Consul"
	"Packages/src/pkg/KVP/KVP"
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
	filter := Filter.GetFilter()
	consulClient := consul.Init("consul:8500")

	//First run settings
	if config.AppFirstRun {
		repo.CreateInitialData(30)
		configKVP := KVP.GetKVPs(config, "/", config.ApplicationName, map[string]string{})
		for k, v := range configKVP {
			_ = consulClient.Set(k, v)
		}
	}

	Middlewares.UsePanicHandlerMiddleware(e)
	Handlers.NewHandler(e, repo, filter, health, &consulClient)
	e.Start(":80")
}
