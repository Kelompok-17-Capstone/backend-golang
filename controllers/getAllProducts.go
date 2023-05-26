package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
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

// get product by ID

func GetProductByIDController(c echo.Context) error {
	id := c.Param("id")
	product, err := usecase.GetProductByid(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item",
		"data":    product,
	})

}
