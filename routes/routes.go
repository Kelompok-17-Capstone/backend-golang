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

	p := e.Group("profile", jwt.JWT([]byte(constants.SECRET_KEY)))
	p.GET("", controllers.ViewMemberInformationController)
	p.POST("", controllers.CreateUserProfileController)
	p.PUT("", controllers.RegisterAsMemberController)
	p.PUT("/photo", controllers.UpdateUserPhotoController)

	products := e.Group("admin/products", jwt.JWT([]byte(constants.SECRET_KEY)))
	products.GET("", controllers.GetProductsController)
	products.POST("", controllers.CreateProductController)
	products.GET("/:id", controllers.GetProductByIDController)
	products.DELETE("/:id", controllers.DeleteProductController)
	products.PUT("/:id", controllers.UpdateProductController)

	users := e.Group("admin/users", jwt.JWT([]byte(constants.SECRET_KEY)))
	users.GET("", controllers.GetUsersController)
	users.GET("/:id", controllers.GetUserController)

	return e
}
