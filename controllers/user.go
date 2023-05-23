package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserProfileController(c echo.Context) error {
	var req payload.Profile
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	id := middlewares.GetUserLoginId(c)
	if err := usecase.CreateUserProfil(id, &req); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create profil",
	})
}
