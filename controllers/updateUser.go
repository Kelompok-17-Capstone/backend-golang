package controllers

import (
	"backend-golang/models/payload"
	"backend-golang/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := payload.UpdateUser{}
	c.Bind(&req)
	if err := usecase.UpdateUser(&req, uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
