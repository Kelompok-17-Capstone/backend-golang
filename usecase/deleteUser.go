package usecase

import (
	"backend-golang/repository/database"
	"fmt"
)

func DeleteUser(id uint) (err error) {
	if _, err := database.GetUser(id); err != nil {
		return err
	}
	err = database.DeleteUser(id)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err: ", err)
		return
	}
	return
}
