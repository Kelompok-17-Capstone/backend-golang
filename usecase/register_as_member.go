package usecase

import (
	"backend-golang/repository/database"
)

func RegisterAsMember(id uint) error {
	user, err := database.ViewMemberInformation(id)
	if err != nil {
		return err
	}

	user.Role = "member"
	err = database.UpdateUserRole(user.ID, user.Role)
	if err != nil {
		return err
	}

	memberCode := database.GenerateMemberCode(user.ID)
	err = database.StoreMemberCode(user.ID, memberCode)
	if err != nil {
		return err
	}

	return nil
}
