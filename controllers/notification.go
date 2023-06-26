package controllers

import (
	"backend-golang/middlewares"
	"backend-golang/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetNotifications(c echo.Context) error {
	userId := middlewares.GetUserLoginId(c)
	resp, err := usecase.GetNotifications(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to get notifications",
		"data":    resp,
	})
}

func GetNotificationById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := usecase.GetNotificationById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to get notification",
		"data":    resp,
	})
}

func UpdateNotificationStatus(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := usecase.UpdateNotificationStatus(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to update status notification",
	})
}
func DeleteNotification(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := usecase.DeleteNotification(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to delete notification",
	})
}
