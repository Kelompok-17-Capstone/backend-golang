package routes

import (
	"backend-golang/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.RegisterController)

	return e
}
