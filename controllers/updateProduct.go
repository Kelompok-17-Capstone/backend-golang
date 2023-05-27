package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func UpdateProductController(c echo.Context) error {
	var product payload.UpdateProduct
	id := c.Param("id")
	c.Bind(&product)
	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	if err := usecase.UpdateProduct(uuid.FromStringOrNil(id), &product, image); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update product",
	})
}
