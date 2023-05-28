package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func GetUsers(keyword, role string) (response []payload.GetUser, err error) {
	var users []models.User
	if keyword == "" {
		if role == "" {
			u, e := database.GetUsers()
			users = u
			err = e
		} else if role == "reguler" {
			u, e := database.GetUsersByRole("reguler")
			users = u
			err = e
		} else if role == "member" {
			u, e := database.GetUsersByRole("member")
			users = u
			err = e
		}
	} else {
		if role == "" {
			p, e := database.GetUsersByName(keyword)
			users = p
			err = e
		} else if role == "reguler" {
			p, e := database.GetUsersByNameAndRole(keyword, "reguler")
			users = p
			err = e
		} else if role == "member" {
			p, e := database.GetUsersByNameAndRole(keyword, "member")
			users = p
			err = e
		}
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
