package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProdcutController(c echo.Context) error {
	var req payload.CreateProduct
	c.Bind(&req)
	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := usecase.CreateProduct(&req, image); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create product",
	})
}
