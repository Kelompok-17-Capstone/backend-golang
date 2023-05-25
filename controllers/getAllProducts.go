package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all product

func GetAllProductController(c echo.Context) error {
	var products []payload.ProductResponse
	products, err := usecase.GetAllProduct()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all product",
		"products": products,
	})
}
