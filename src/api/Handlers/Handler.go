package Handlers

import (
	UserRepository "Packages/src/api/Repository"
	"Packages/src/api/Type/ErrorTypes"
	"Packages/src/pkg/HealthChecks"
	"Packages/src/pkg/Logger"
	"Packages/src/pkg/MongoFilter"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Repository   UserRepository.IRepository
	Logger       *Logger.Logger
	Filter       *MongoFilter.MongoFilter
	HealthChecks *HealthChecks.ServiceHealth
}

func NewHandler(e *echo.Echo, repository UserRepository.IRepository, logger *Logger.Logger, filter *MongoFilter.MongoFilter, hc *HealthChecks.ServiceHealth) *Handler {
	handler := &Handler{
		Repository:   repository,
		Logger:       logger,
		Filter:       filter,
		HealthChecks: hc,
	}
	g := e.Group("/users")
	g.GET("", handler.GetMany)
	g.POST("", handler.Create)
	g.GET("/:id", handler.GetSingle)
	e.GET("/quickhealth", handler.QuickHealth)
	e.GET("/health", handler.Health)
	e.GET("/err", handler.ThrowPanic)
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
