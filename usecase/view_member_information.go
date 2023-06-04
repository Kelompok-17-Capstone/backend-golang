package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"

	"strings"
)

func ViewMemberInformation(id uint) (payload.GetMember, error) {
	user, err := database.ViewMemberInformation(id)
	if err != nil {
		return payload.GetMember{}, err
	}

	words := strings.Fields(user.Profile.Name)
	var username string
	if len(words) >= 8 {
		username = strings.Join(words[:8], " ")
	} else {
		username = user.Profile.Name
	}

	memberCode, err := database.GetMemberCode(user.ID)
	if err != nil {
		return payload.GetMember{}, err
	}
	address := user.Profile.Address.Address + ", " + user.Profile.Address.City + ", " + user.Profile.Address.Province
	resp := payload.GetMember{
		ID:          user.ID,
		Name:        username,
		Email:       user.Email,
		PhoneNumber: user.Profile.PhoneNumber,
		Address:     address,
		Image:       user.Profile.Photo,
		MemberCode:  memberCode,
	}
	return resp, nil
}
