package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateDetailCartItemController(c echo.Context) error {
	var cart *payload.UpdateQtyCart
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	c.Bind(&cart)
	if err := c.Validate(cart); err != nil {
		return err
	}
	if err := usecase.UpdateCartQty(uint(id), cart); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update quantity",
	})
}
