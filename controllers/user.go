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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create profil",
	})
}

func UpdateUserPhotoController(c echo.Context) error {

	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to upload photo")
	}
	id := middlewares.GetUserLoginId(c)
	if err := usecase.UploadPhotoProfile(id, file); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update photo profile",
	})
}
