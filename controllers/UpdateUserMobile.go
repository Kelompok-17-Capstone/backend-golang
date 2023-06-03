package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// update user email
func UpdateUserEmailController(c echo.Context) error {
	var req payload.UpdateEmail
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	id := middlewares.GetUserLoginId(c)
	if err := usecase.UpdateUserEmail(id, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update email",
	})
}

// update user password
func UpdatePasswordController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	var req payload.UpdatePassword
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	user, err := usecase.UpdatePassword(userID, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// update name user
func UpdateNameController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	var req payload.Profile
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := usecase.UpdateName(userID, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Name updated successfully",
	})
}

// update phone number
func UpdatePhoneNumberController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	var req payload.Profile
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := usecase.UpdatePhoneNumber(userID, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Phone number updated successfully",
	})
}

// update address
func UpdateAddressController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)

	var req payload.UpdateAddress
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := usecase.UpdateAddress(userID, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Address updated successfully",
	})
}
