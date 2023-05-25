package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func ViewMemberInformation(id uint) (payload.GetMember, error) {
	user, err := database.ViewMemberInformation(id)
	if err != nil {
		return payload.GetMember{}, err
	}
	resp := payload.GetMember{
		ID:    user.ID,
		Name:  user.Profile.Name,
		Image: user.Profile.Photo,
	}
	return resp, nil
}
