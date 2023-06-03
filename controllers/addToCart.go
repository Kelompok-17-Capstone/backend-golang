package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddToCartController(c echo.Context) error {
	// Get userID from authenticated user
	userID := middlewares.GetUserLoginId(c)
	// Parse request body
	req := new(payload.AddToCart)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Validate request
	if err := c.Validate(req); err != nil {
		return err
	}

	// Add to cart
	err := usecase.AddToCart(userID, req.ProductID, req.Quantity)
	if err != nil {
		if strings.Contains(err.Error(), "product stock is not available") {
			return echo.NewHTTPError(http.StatusBadRequest, "Product stock is not available")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add to cart")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product added to cart successfully",
	})
}
