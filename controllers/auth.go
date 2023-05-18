package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	var req payload.Register
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.Register(&req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
	})
}
