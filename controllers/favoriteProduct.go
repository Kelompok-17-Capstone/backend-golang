package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddFavoriteProductController(c echo.Context) error {
	var req payload.AddFavoriteProduct
	userID := middlewares.GetUserLoginId(c)
	c.Bind(&req)

	if err := usecase.AddFavoriteProduct(userID, req.ProductID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, "Favorite product added successfully")
}

func GetFavoriteProductController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	favoriteProduct, err := usecase.GetFavoriteProduct(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get favorite products",
		"products": favoriteProduct,
	})
}

func DeleteFavoriteProductController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)
	favID, _ := strconv.Atoi(c.Param("id"))
	err := usecase.DeleteFavoriteProduct(userID, uint(favID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Favorite product deleted successfully",
	})
}
