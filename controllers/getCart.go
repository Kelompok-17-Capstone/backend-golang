package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCartController(c echo.Context) error {
	// Get userID from authenticated user
	userID := middlewares.GetUserLoginId(c)

	// Get cart
	resp, err := usecase.GetCart(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get cart")
	}

	// Prepare response payload
	response := map[string]interface{}{
		"message": "Cart retrieved successfully",
		"data":    resp,
	}

	return c.JSON(http.StatusOK, response)
}
