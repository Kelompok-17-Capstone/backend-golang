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
	e.Use(middleware.CORS())

	e.GET("/guest", controllers.Guest)
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	e.GET("/products", controllers.GetProductsMobileController)
	e.GET("/products/:id", controllers.GetProductMobileByIdController)
	user := e.Group("", jwt.JWT([]byte(constants.SECRET_KEY)), m.IsUser)
	user.GET("/home", controllers.Home)
	n := user.Group("notifications")
	n.GET("", controllers.GetNotifications)
	n.GET("/:id", controllers.GetNotificationById)
	n.PUT("/status/:id", controllers.UpdateNotificationStatus)
	n.DELETE("/:id", controllers.DeleteNotification)
	p := user.Group("profile", jwt.JWT([]byte(constants.SECRET_KEY)))
	p.GET("", controllers.ViewMemberInformationController)
	p.POST("", controllers.CreateUserProfileController)
	p.POST("/address", controllers.CreateAddressController)
	p.PUT("/email", controllers.UpdateUserEmailController)
	p.PUT("/password", controllers.UpdatePasswordController)
	p.PUT("/name", controllers.UpdateNameController)
	p.PUT("/phone-number", controllers.UpdatePhoneNumberController)
	p.PUT("/address/:id", controllers.UpdateAddressController)
	p.DELETE("/address/:id", controllers.DeleteAddressController)
	p.PUT("", controllers.RegisterAsMemberController)
	p.PUT("/photo", controllers.UpdateUserPhotoController)
	//topup
	tp := user.Group("topup")
	tp.POST("", controllers.CreateTopupController)
	coin := user.Group("coin")
	coin.GET("", controllers.GetCoinController)
	balance := user.Group("balance")
	balance.GET("", controllers.GetBalanceController)
	//get order mobile
	orderMobile := user.Group("orders")
	orderMobile.GET("", controllers.GetOrderMobileController)
	orderMobile.POST("", controllers.CreateOrderController)
	//cart
	cart := user.Group("/cart")
	cart.POST("", controllers.AddToCartController)
	cart.GET("", controllers.GetCartController)
	cart.DELETE("/:id", controllers.DeleteDetailCartItemController)
	cart.PUT("/:id", controllers.UpdateDetailCartItemController)
	favoriteProducts := user.Group("favorite")
	favoriteProducts.GET("", controllers.GetFavoriteProductController)
	favoriteProducts.POST("", controllers.AddFavoriteProductController)
	favoriteProducts.DELETE("/:id", controllers.DeleteFavoriteProductController)

	admin := e.Group("admin", jwt.JWT([]byte(constants.SECRET_KEY)), m.IsAdmin)
	dashboard := admin.Group("/dashboard")
	dashboard.GET("", controllers.GetDashboardController)
	products := admin.Group("/products")
	products.GET("", controllers.GetProductsController)
	products.POST("", controllers.CreateProductController)
	products.GET("/:id", controllers.GetProductByIDController)
	products.PUT("/:id", controllers.UpdateProductController)
	products.DELETE("/:id", controllers.DeleteProductController)
	users := admin.Group("/users")
	users.GET("", controllers.GetUsersController)
	users.GET("/:id", controllers.GetUserController)
	users.PUT("/:id", controllers.UpdateUserController)
	users.DELETE("/:id", controllers.DeleteUserController)
	orders := admin.Group("/orders", jwt.JWT([]byte(constants.SECRET_KEY)), m.IsAdmin)
	orders.GET("", controllers.GetOrdersController)
	orders.GET("/:id", controllers.GetOrderController)
	orders.PUT("/:id", controllers.UpdateOrderController)
	orders.DELETE("/:id", controllers.DeleteOrderController)

	return e
}
