package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetBalanceController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	resp, err := usecase.GetBalance(userID)
	if err != nil {
		if strings.Contains(err.Error(), "Balance record not found") {
			return echo.NewHTTPError(http.StatusBadRequest, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get balance")
	}

	return c.JSON(http.StatusOK, resp)
}
