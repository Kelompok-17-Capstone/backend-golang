package controllers

import (
	"backend-golang/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := usecase.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
