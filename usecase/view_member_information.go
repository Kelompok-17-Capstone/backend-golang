package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"

	"strings"
)

func ViewMemberInformation(id uint) (payload.GetMemberMobile, error) {
	user, err := database.ViewMemberInformation(id)
	if err != nil {
		return payload.GetMemberMobile{}, err
	}

	words := strings.Fields(user.Profile.Name)
	var username string
	if len(words) >= 8 {
		username = strings.Join(words[:8], " ")
	} else {
		username = user.Profile.Name
	}

	var address []payload.Address
	for _, value := range user.Profile.Address {
		address = append(address, payload.Address{
			ID:      value.ID,
			Address: value.Address + ", " + value.City + ", " + value.Province,
			Status:  value.Status,
		})
	}

	resp := payload.GetMemberMobile{
		ID:          user.ID,
		Name:        username,
		Email:       user.Email,
		PhoneNumber: user.Profile.PhoneNumber,
		Address:     address,
		Image:       user.Profile.Photo,
		MemberCode:  user.MemberCode,
		Balance:     user.Balance,
		Coin:        user.Coin,
		Barcode:     user.Barcode,
	}
	return resp, nil
}
