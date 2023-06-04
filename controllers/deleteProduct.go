package controllers

import (
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func DeleteProductController(c echo.Context) error {
	id := c.Param("id")
	if err := usecase.DeleteProduct(uuid.FromStringOrNil(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Product",
	})
}
