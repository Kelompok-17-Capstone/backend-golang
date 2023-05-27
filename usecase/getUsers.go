package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func GetUsers() (response []payload.GetUser, err error) {
	users, err := database.GetUsers()
	if err != nil {
		return
	}
	for _, user := range users {
		response = append(response, payload.GetUser{
			Name:        user.Profile.Name,
			Email:       user.Email,
			PhoneNumber: user.Profile.PhoneNumber,
			Address:     user.Profile.Address.Province,
		})
	}
	return
}

func GetUser(id uint) (resp payload.GetUser, err error) {
	user, err := database.GetUser(id)
	if err != nil {
		return payload.GetUser{}, err
	}
	resp = payload.GetUser{
		Name:        user.Profile.Name,
		Email:       user.Email,
		PhoneNumber: user.Profile.PhoneNumber,
		Address:     user.Profile.Address.Province,
	}
	return
}
