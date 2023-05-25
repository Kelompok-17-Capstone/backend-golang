package routes

import (
	"backend-golang/constants"
	"backend-golang/controllers"
	m "backend-golang/middlewares"

	jwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LoggerMiddleware(e)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	e.POST("/profil", controllers.CreateUserProfileController, jwt.JWT([]byte(constants.SECRET_KEY)))

	products := e.Group("admin/products", jwt.JWT([]byte(constants.SECRET_KEY)))
	products.POST("", controllers.CreateProdcutController)

	return e
}
