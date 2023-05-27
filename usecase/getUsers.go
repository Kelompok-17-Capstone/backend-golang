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
		address := user.Profile.Address.Address + ", " + user.Profile.Address.City + ", " + user.Profile.Address.Province
		response = append(response, payload.GetUser{
			ID:          user.ID,
			Name:        user.Profile.Name,
			Email:       user.Email,
			PhoneNumber: user.Profile.PhoneNumber,
			Address:     address,
		})
	}
	return
}

func GetUser(id uint) (resp payload.GetUser, err error) {
	user, err := database.GetUser(id)
	if err != nil {
		return payload.GetUser{}, err
	}
	address := user.Profile.Address.Address + ", " + user.Profile.Address.City + ", " + user.Profile.Address.Province
	resp = payload.GetUser{
		ID:          user.ID,
		Name:        user.Profile.Name,
		Email:       user.Email,
		PhoneNumber: user.Profile.PhoneNumber,
		Address:     address,
	}
	return
}
