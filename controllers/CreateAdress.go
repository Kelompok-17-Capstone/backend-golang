package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// create address
func CreateAddressController(c echo.Context) error {
	var req payload.CreateAddress
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	id := middlewares.GetUserLoginId(c)
	if err := usecase.CreateAddress(id, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create address",
	})
}

// func CreateAddressController(c echo.Context) error {
// 	userID := middlewares.GetUserLoginId(c)

// 	var req payload.CreateAddress
// 	if err := c.Bind(&req); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
// 	}

// 	// if err := c.Validate(&req); err != nil {
// 	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	// }

// 	if err := usecase.CreateAddress(userID, &req); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Address created successfully",
// 	})
// }
