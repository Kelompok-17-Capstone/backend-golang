package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// get all product

func GetProductsController(c echo.Context) error {
	var products []payload.ProductResponse
	status := c.QueryParam("status")
	keyword := c.QueryParam("keyword")
	products, err := usecase.GetProducts(status, keyword)
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
		"message": "success get products by id",
		"data":    product,
	})

}

func GetProductsMobileController(c echo.Context) error {
	tab := c.QueryParam("tab")
	keyword := c.QueryParam("keyword")
	price := c.QueryParam("price")
	products, err := usecase.GetProductsMobile(keyword, tab, price)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all product",
		"products": products,
	})
}

func GetProductMobileByIdController(c echo.Context) error {
	id := c.Param("id")
	product, err := usecase.GetProductMobileByid(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get products by id",
		"data":    product,
	})

}
