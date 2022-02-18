package Middlewares

import (
	"Packages/src/api/Logging"
	"Packages/src/api/Type/ErrorTypes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PanicHandlerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := Logging.GetLogger().CreateLog().InformRequest(c.Request())
			defer func() {
				if err := recover(); err != nil {
					switch v := err.(type) {
					case *ErrorTypes.Error:
						log.AddStruct(*v)
						c.JSON(v.StatusCode, v.PublicError)
					case ErrorTypes.Error:
						log.AddStruct(v)
						c.JSON(v.StatusCode, v.PublicError)
					default:
						log.AddStruct(v)
						c.JSON(http.StatusInternalServerError, err)
					}
				}
				log.InformResponse(c.Response().Status, c.Response().Size)
				go log.LogWithRules()
			}()
			return next(c)
		}
	}
}

func UsePanicHandlerMiddleware(e *echo.Echo) {
	e.Use(PanicHandlerMiddleware())
}
