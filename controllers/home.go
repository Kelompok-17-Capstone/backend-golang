package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	userId := middlewares.GetUserLoginId(c)
	home, err := usecase.Home(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success get home",
		"data":    home,
	})
}

func Guest(c echo.Context) error {
	home, err := usecase.Guest()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success get home",
		"data":    home,
	})
}
