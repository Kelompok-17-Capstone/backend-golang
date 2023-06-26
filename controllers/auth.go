package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	var req payload.Register
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := usecase.Register(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := middlewares.CreateToken(user.Email, user.Role, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
		"token":   token,
	})
}

func LoginController(c echo.Context) error {
	var req payload.Login
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := usecase.Login(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := middlewares.CreateToken(user.Email, user.Role, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success login user",
		"token":       token,
		"status_user": user.Status,
	})
}
