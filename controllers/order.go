package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateOrderController(c echo.Context) error {
	var req payload.CreateOrder
	userID := middlewares.GetUserLoginId(c)
	c.Bind(&req)
	if err := usecase.CreateOrder(userID, &req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Order product success",
	})
}

func GetOrdersController(c echo.Context) error {
	req := payload.OrdersParam{
		Keyword: c.QueryParam("keyword"),
		Status:  c.QueryParam("status"),
	}
	orders, err := usecase.GetOrders(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Get orders success",
		"orders":  orders,
	})
}
