package usecase

import (
	"backend-golang/repository/database"
)

func RegisterAsMember(id uint) error {
	err := database.UpdateUserRole(id, "member")
	if err != nil {
		return err
	}

	user, err := database.GetUserProfile(id)
	if err != nil {
		return err
	}

	user.Role = "member"
	err = database.UpdateUserRole(user.ID, user.Role)
	if err != nil {
		return err
	}

	return nil
}
