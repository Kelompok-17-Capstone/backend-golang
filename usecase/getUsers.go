package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func GetUsers(keyword, role string) (response []payload.GetUser, err error) {
	req := payload.UsersParam{
		Keyword: keyword,
		Role:    role,
	}
	users, err := database.GetUsers(&req)

	for _, user := range users {
		var address string
		if len(user.Profile.Address) > 0 {
			address = user.Profile.Address[0].Address + ", " + user.Profile.Address[0].City + ", " + user.Profile.Address[0].Province
		}
		response = append(response, payload.GetUser{
			ID:          user.ID,
			Name:        user.Profile.Name,
			Email:       user.Email,
			PhoneNumber: user.Profile.PhoneNumber,
			Address:     address,
			Status:      user.Role,
		})
	}
	return
}

func GetUser(id uint) (resp payload.GetUser, err error) {
	user, err := database.GetUser(id)
	if err != nil {
		return payload.GetUser{}, err
	}
	var address string
	if len(user.Profile.Address) > 0 {
		address = user.Profile.Address[0].Address + ", " + user.Profile.Address[0].City + ", " + user.Profile.Address[0].Province
	}
	resp = payload.GetUser{
		ID:          user.ID,
		Name:        user.Profile.Name,
		Email:       user.Email,
		PhoneNumber: user.Profile.PhoneNumber,
		Address:     address,
		Status:      user.Role,
	}
	return
}
