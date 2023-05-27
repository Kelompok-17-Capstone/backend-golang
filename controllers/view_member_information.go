package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ViewMemberInformationController(c echo.Context) error {
	id := middlewares.GetUserLoginId(c)
	user, err := usecase.ViewMemberInformation(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
