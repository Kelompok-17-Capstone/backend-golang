package controllers

import (
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDashboardController(c echo.Context) error {
	resp, err := usecase.GetDashboard()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Something went wrong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get dashboard",
		"data":    resp,
	})
}
