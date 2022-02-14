package HealthChecks

import (
	"Packages/src/Configs"
	"Packages/src/pkg/HealthChecks"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetHealthChecks(client *mongo.Client, configs *Configs.AppConfig) *HealthChecks.ServiceHealth {
	hc := HealthChecks.NewServiceHealth()
	hc.AddHealthCheck("Mongo",
		"Database, DB, Mongo",
		"Mongo User Database",
		true,
		time.Second*5, MongoHealthCheck(client, configs))
	return hc
}
