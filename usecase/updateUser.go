package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
	"fmt"
)

func UpdateUser(user *models.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}

	return
}
