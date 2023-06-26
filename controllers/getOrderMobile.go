package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetOrderMobileController(c echo.Context) error {
	userID := middlewares.GetUserLoginId(c)
	req := payload.GetOrderMobileParam{
		Status: c.QueryParam("status"),
	}
	resp, err := usecase.GetOrderMobile(userID, req.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	response := map[string]interface{}{
		"message": "Orders retrieved successfully",
		"data":    resp,
	}

	return c.JSON(http.StatusOK, response)
}
