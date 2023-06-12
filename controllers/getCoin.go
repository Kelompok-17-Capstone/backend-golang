package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetCoinController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	resp, err := usecase.GetCoin(userID)
	if err != nil {
		if strings.Contains(err.Error(), "Coin record not found") {
			return echo.NewHTTPError(http.StatusBadRequest, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get coin")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get coin history success",
		"history": resp,
	})
}
