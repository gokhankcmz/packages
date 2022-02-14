package main

import (
	Configs "Packages/src/Configs"
	"Packages/src/api/Filter"
	"Packages/src/api/Handlers"
	"Packages/src/api/HealthChecks"
	Logging "Packages/src/api/Logger"
	"Packages/src/api/Middlewares"
	"Packages/src/api/Mongo"
	UserRepository "Packages/src/api/Repository"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Application Starting...")
	e := echo.New()
	Configs.SetDefault()

	config := Configs.GetConfigs()
	logger := Logging.GetLogger(config)
	filter := Filter.GetFilter()

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

	if config.AppFirstRun {
		repo.CreateInitialData(30)
	}

	Middlewares.UsePanicHandlerMiddleware(e, logger)
	Handlers.NewHandler(e, repo, logger, filter, health)
	e.Start(":80")
}
