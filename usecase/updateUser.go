package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func UpdateUser(req *payload.UpdateUser, id uint) (err error) {
	if _, err := database.GetUser(id); err != nil {
		return err
	}
	profile := models.Profile{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	if err := database.UpdateUserProfil(&profile, id); err != nil {
		return err
	}
	profil, err := database.GetUserProfile(id)
	if err != nil {
		return err
	}
	address := models.Address{
		Address:  req.Address,
		City:     req.City,
		Province: req.Province,
	}
	if err := database.UpdateAddressWhereStatus(profil.ID, &address, true); err != nil {
		return err
	}
	if err := database.UpdateUserRole(id, req.Status); err != nil {
		return err
	}

	return
}
