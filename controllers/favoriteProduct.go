package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddFavoriteProductController(c echo.Context) error {
	userID := c.Get("userID").(uint)

	var addFavoriteProduct payload.AddFavoriteProduct
	if err := c.Bind(&addFavoriteProduct); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	productID := addFavoriteProduct.ProductID.String()

	err := usecase.AddFavoriteProduct(userID, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to add favorite product")
	}

	return c.JSON(http.StatusCreated, "Favorite product added successfully")
}

func GetFavoriteProductController(c echo.Context) error {
	userID := c.Get("userID").(uint)

	favoriteProduct, err := usecase.GetFavoriteProduct(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get favorite product")
	}

	return c.JSON(http.StatusOK, favoriteProduct)
}

func DeleteFavoriteProductController(c echo.Context) error {
	userID := c.Get("userID").(uint)

	var deleteFavoriteProduct payload.DeleteFavoriteProduct
	if err := c.Bind(&deleteFavoriteProduct); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	productID := deleteFavoriteProduct.ProductID.String()

	err := usecase.DeleteFavoriteProduct(userID, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete favorite product")
	}

	return c.JSON(http.StatusOK, "Favorite product deleted successfully")
}
