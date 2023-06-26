package controllers

import (
	"backend-golang/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get  all user

func GetUsersController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	role := c.QueryParam("role")
	users, err := usecase.GetUsers(keyword, role)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := usecase.GetUser(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}
