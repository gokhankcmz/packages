package HealthChecks

import (
	"Packages/src/Configs"
	"Packages/src/pkg/HealthChecks"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func MongoHealthCheck(client *mongo.Client, configs *Configs.AppConfig) func() *HealthChecks.HealthCheckResult {
	return func() *HealthChecks.HealthCheckResult {
		ctx, _ := context.WithTimeout(context.Background(), configs.MongoConnectionDuration)
		if ctx == nil {
			ctx = context.Background()
		}
		db := client.Database(configs.DBName)
		startTime := time.Now()
		res := db.RunCommand(ctx, bson.D{
			{"ping", 1},
		}, nil)

		if res.Err() != nil {
			return &HealthChecks.HealthCheckResult{
				Status:   HealthChecks.Unhealthy,
				Clusters: nil,
				Time:     time.Now(),
				Duration: time.Since(startTime),
			}
		}
		duration := time.Since(startTime)
		if duration.Milliseconds() > 500 {
			return &HealthChecks.HealthCheckResult{
				Status:   HealthChecks.Degraded,
				Clusters: nil,
				Time:     time.Now(),
				Duration: duration,
			}
		}
		return &HealthChecks.HealthCheckResult{
			Status:   HealthChecks.Healthy,
			Clusters: nil,
			Time:     time.Now(),
			Duration: duration,
		}
	}

}
