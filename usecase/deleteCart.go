package usecase

import (
	"backend-golang/repository/database"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func DeleteDetailCartItem(id uint) error {
	_, err := database.GetDetailCartItemById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("detail cart item not found")
		}
		return err
	}
	err = database.DeleteDetailCartItem(id)
	if err != nil {
		fmt.Println("Delete: error deleting item , err:", err)
		return err
	}
	return nil
}
