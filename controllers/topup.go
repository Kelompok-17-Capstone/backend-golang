package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateTopupController(c echo.Context) error {
	// Get userID from authenticated user
	userID := middlewares.GetUserLoginId(c)
	// Parse request body
	req := new(payload.BalanceRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Validate request
	if err := c.Validate(req); err != nil {
		return err
	}

	err := usecase.CreateTopup(userID, req.Total)
	if err != nil {
		if strings.Contains(err.Error(), "cant create topup") {
			return echo.NewHTTPError(http.StatusBadRequest, "cant create topup")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add balance")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "balance added successfully",
	})
}
