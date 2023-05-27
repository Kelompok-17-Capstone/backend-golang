package usecase

import (
	"backend-golang/repository/database"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func DeleteProduct(id uuid.UUID) (err error) {
	if _, err := database.GetProductById(id); err != nil {
		return err
	}
	err = database.DeleteProduct(id)
	if err != nil {
		fmt.Println("Delete: error deleting Product, err:", err)
		return err
	}
	return nil
}
