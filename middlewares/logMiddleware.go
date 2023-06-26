package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time = ${time_rfc3339}, host = ${host}, method = ${method}, uri = ${uri}, status = ${status}\n",
	}))
}
