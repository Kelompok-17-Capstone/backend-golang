package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterAsMemberController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	err := usecase.RegisterAsMember(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully registered as a member",
	})
}
