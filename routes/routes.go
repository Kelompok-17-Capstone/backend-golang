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
	e.GET("/products", controllers.GetProductsMobileController)

	p := e.Group("profile", jwt.JWT([]byte(constants.SECRET_KEY)))
	p.GET("", controllers.ViewMemberInformationController)
	p.POST("", controllers.CreateUserProfileController)
	p.POST("/address", controllers.CreateAddressController)
	p.PUT("/email", controllers.UpdateUserEmailController)
	p.PUT("/password", controllers.UpdatePasswordController)
	p.PUT("/name", controllers.UpdateNameController)
	p.PUT("/phone-number", controllers.UpdatePhoneNumberController)
	p.PUT("/address", controllers.UpdateAddressController)
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
	users.PUT("/:id", controllers.UpdateUserController)
	users.DELETE("/:id", controllers.DeleteUserController)

	m := e.Group("member", jwt.JWT([]byte(constants.SECRET_KEY)))
	m.POST("", controllers.RegisterAsMemberController)
	m.GET("", controllers.ViewMemberInformationController)

	//cart
	cart := e.Group("/cart", jwt.JWT([]byte(constants.SECRET_KEY)))
	cart.POST("", controllers.AddToCartController)
	cart.GET("", controllers.GetCartController)
	cart.DELETE("/:id", controllers.DeleteDetailCartItemController)
	cart.PUT("/:id", controllers.UpdateDetailCartItemController)

	favoriteProducts := e.Group("favorite", jwt.JWT([]byte(constants.SECRET_KEY)))
	favoriteProducts.GET("", controllers.GetFavoriteProductController)
	favoriteProducts.POST("", controllers.AddFavoriteProductController)
	favoriteProducts.DELETE("/:id", controllers.DeleteFavoriteProductController)

	return e
}
