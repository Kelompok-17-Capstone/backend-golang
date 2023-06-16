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

	e.GET("/home", controllers.Home, jwt.JWT([]byte(constants.SECRET_KEY)))
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	e.POST("/orders", controllers.CreateOrderController, jwt.JWT([]byte(constants.SECRET_KEY)))
	e.GET("/products", controllers.GetProductsMobileController)
	e.GET("/products/:id", controllers.GetProductMobileByIdController)

	p := e.Group("profile", jwt.JWT([]byte(constants.SECRET_KEY)))
	p.GET("", controllers.ViewMemberInformationController)
	p.POST("", controllers.CreateUserProfileController)
	p.POST("/address", controllers.CreateAddressController)
	p.PUT("/email", controllers.UpdateUserEmailController)
	p.PUT("/password", controllers.UpdatePasswordController)
	p.PUT("/name", controllers.UpdateNameController)
	p.PUT("/phone-number", controllers.UpdatePhoneNumberController)
	p.PUT("/address/:id", controllers.UpdateAddressController)
	p.PUT("", controllers.RegisterAsMemberController)
	p.PUT("/photo", controllers.UpdateUserPhotoController)

	products := e.Group("admin/products", jwt.JWT([]byte(constants.SECRET_KEY)))
	products.GET("", controllers.GetProductsController)
	products.POST("", controllers.CreateProductController)
	products.GET("/:id", controllers.GetProductByIDController)
	products.PUT("/:id", controllers.UpdateProductController)
	products.DELETE("/:id", controllers.DeleteProductController)

	users := e.Group("admin/users", jwt.JWT([]byte(constants.SECRET_KEY)))
	users.GET("", controllers.GetUsersController)
	users.GET("/:id", controllers.GetUserController)
	users.PUT("/:id", controllers.UpdateUserController)
	users.DELETE("/:id", controllers.DeleteUserController)

	orders := e.Group("admin/orders", jwt.JWT([]byte(constants.SECRET_KEY)))
	orders.GET("", controllers.GetOrdersController)
	orders.GET("/:id", controllers.GetOrderController)
	orders.PUT("/:id", controllers.UpdateOrderController)
	orders.DELETE("/:id", controllers.DeleteOrderController)

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

	//topup
	tp := e.Group("topup", jwt.JWT([]byte(constants.SECRET_KEY)))
	tp.POST("", controllers.CreateTopupController)
	// tp.GET("", controllers.)

	coin := e.Group("coin", jwt.JWT([]byte(constants.SECRET_KEY)))
	coin.GET("", controllers.GetCoinController)

	balance := e.Group("balance", jwt.JWT([]byte(constants.SECRET_KEY)))
	balance.GET("", controllers.GetBalanceController)

	return e
}
