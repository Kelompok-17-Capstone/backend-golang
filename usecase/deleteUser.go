package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
	"fmt"
)

func DeleteUser(id uint) (err error) {
	user := models.User{}
	user.ID = id
	err = database.DeleteUser(id)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err: ", err)
		return
	}

	return
}
