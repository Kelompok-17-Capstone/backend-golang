package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
)

func ViewMemberInformation(id uint) (models.User, error) {
	user, err := database.ViewMemberInformation(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
