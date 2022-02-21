package Handlers

import (
	"Packages/src/Configs"
	"Packages/src/api/Filter"
	"Packages/src/api/Logging"
	"Packages/src/api/Middlewares"
	UserRepository "Packages/src/api/Repository"
	"Packages/src/api/Type/ErrorTypes"
	consul "Packages/src/pkg/Consul"
	"Packages/src/pkg/HealthChecks"
	KvpConverter "Packages/src/pkg/KVP"
	"Packages/src/pkg/MongoFilter"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Repository   UserRepository.IRepository
	HealthChecks *HealthChecks.ServiceHealth
	Consul       *consul.Client
	Filter       *MongoFilter.MongoFilter
	e            *echo.Echo
	KvpConverter *KvpConverter.Settings
}

func NewHandler(e *echo.Echo, repository UserRepository.IRepository, filter *MongoFilter.MongoFilter, hc *HealthChecks.ServiceHealth, client *consul.Client, kvpConverter *KvpConverter.Settings) *Handler {
	handler := &Handler{
		Repository:   repository,
		HealthChecks: hc,
		Filter:       filter,
		Consul:       client,
		e:            e,
		KvpConverter: kvpConverter,
	}
	handler.SetEndpoints()
	Middlewares.UsePanicHandlerMiddleware(e)
	e.HideBanner = true
	return handler
}
func (h Handler) SetEndpoints() {
	g := h.e.Group("/users")
	g.GET("", h.GetMany)
	g.POST("", h.Create)
	g.GET("/:id", h.GetSingle)

	h.e.GET("/quickhealth", h.QuickHealth)
	h.e.GET("/health", h.Health)
	h.e.GET("/err", h.ThrowPanic)
	h.e.POST("/watch", h.Watch)
	h.e.GET("/get", h.Get)

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
	config := &Configs.AppConfig{}
	h.KvpConverter.GetObject(config, configKVP)
	Configs.Set(*config)
	Logging.Init(configs)
	h.Filter = Filter.SetDynamicFilter(config)
	h.SetEndpoints()
	return nil
}
func (h Handler) Get(ctx echo.Context) error {
	ctx.JSON(http.StatusOK, Configs.GetConfigs())
	return nil
}
