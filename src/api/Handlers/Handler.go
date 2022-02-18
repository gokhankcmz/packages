package Handlers

import (
	"Packages/src/Configs"
	"Packages/src/api/Logging"
	UserRepository "Packages/src/api/Repository"
	"Packages/src/api/Type/ErrorTypes"
	consul "Packages/src/pkg/Consul"
	"Packages/src/pkg/HealthChecks"
	"Packages/src/pkg/MongoFilter"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Repository   UserRepository.IRepository
	Filter       *MongoFilter.MongoFilter
	HealthChecks *HealthChecks.ServiceHealth
	Consul       *consul.Client
}

func NewHandler(e *echo.Echo, repository UserRepository.IRepository, filter *MongoFilter.MongoFilter, hc *HealthChecks.ServiceHealth, client *consul.Client) *Handler {
	handler := &Handler{
		Repository:   repository,
		Filter:       filter,
		HealthChecks: hc,
		Consul:       client,
	}
	g := e.Group("/users")
	g.GET("", handler.GetMany)
	g.POST("", handler.Create)
	g.GET("/:id", handler.GetSingle)
	e.GET("/quickhealth", handler.QuickHealth)
	e.GET("/health", handler.Health)
	e.GET("/err", handler.ThrowPanic)
	e.POST("/watch", handler.Watch)
	e.GET("/get", handler.Get)
	e.HideBanner = true

	return handler
}

func (h Handler) QuickHealth(ctx echo.Context) error {
	res := h.HealthChecks.GetHealthCheckSummary()
	return ctx.JSON(http.StatusOK, res)
}

func (h Handler) Health(ctx echo.Context) error {
	res := h.HealthChecks.GetHealthCheckResult()
	return ctx.JSON(http.StatusOK, res)
}

func (h Handler) ThrowPanic(ctx echo.Context) error {
	panic(ErrorTypes.InvalidModel)
}

func (h Handler) Watch(ctx echo.Context) error {
	configs := Configs.GetConfigs()
	configKVP := h.Consul.GetConfigs(configs.ApplicationName)
	Configs.SetFromConsul(configs.ApplicationName, configKVP)

	Logging.Init(Configs.GetConfigs())

	return nil
}
func (h Handler) Get(ctx echo.Context) error {
	ctx.JSON(http.StatusOK, Configs.GetConfigs())
	return nil
}
