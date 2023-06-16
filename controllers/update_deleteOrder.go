package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// delete order by id
func DeleteOrderController(c echo.Context) error {
	id := c.Param("id")
	if err := usecase.DeleteOrder(uuid.FromStringOrNil(id)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete order",
	})
}

// update order by ID
func UpdateOrderController(c echo.Context) error {
	var req payload.UpdateOrder
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	id := c.Param("id")
	if err := usecase.UpdateOrderStatusAndArrived(uuid.FromStringOrNil(id), &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update order",
	})
}

// get order by id
func GetOrderController(c echo.Context) error {
	id := c.Param("id")
	order, err := usecase.GetOrderByID(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}
