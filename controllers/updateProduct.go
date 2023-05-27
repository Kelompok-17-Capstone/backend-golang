package controllers

import (
	"backend-golang/models"
	"backend-golang/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func UpdateProductController(c echo.Context) error {
	id := c.Param("id")
	updateProduct := models.Product{}
	c.Bind(&updateProduct)
	updateProduct.ID = uuid.FromStringOrNil(id)
	image, _ := c.FormFile("image")
	if _, err := usecase.UpdateProduct(updateProduct.ID, image); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error updating product",
			"errorDescription": err.Error(),
			"errorMessage":     "Sorry, unable to update product at the moment",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update update Product",
	})
}
