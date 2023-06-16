package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// update status order
func UpdateStatusController(c echo.Context) error {
	var req payload.UpdateOrderStatus
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	id := c.Param("id")
	if err := usecase.UpdateStatusOrder(uuid.FromStringOrNil(id), &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update status",
	})

}

// update arrived at order
func UpdateArrivedAt(c echo.Context) error {
	var req payload.UpadateArrivedAt
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	id := c.Param("id")
	if err := usecase.UpdateArrivedAtOrder(uuid.FromStringOrNil(id), &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update arrived at",
	})

}
