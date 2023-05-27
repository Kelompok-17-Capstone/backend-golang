package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get  all user

func GetUsersController(c echo.Context) error {
	users, err := usecase.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {
	id := middlewares.GetUserLoginId(c)
	user, err := usecase.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}
