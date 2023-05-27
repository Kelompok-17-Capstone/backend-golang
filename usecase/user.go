package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"backend-golang/util"
	"mime/multipart"
)

func CreateUserProfil(id uint, req *payload.Profile) error {
	profile := models.Profile{
		UserID:      id,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	if err := database.CreateUserProfil(&profile); err != nil {
		return err
	}
	address := models.Address{
		ProfileID: profile.ID,
		Address:   req.Address,
		City:      req.City,
		Province:  req.Province,
	}
	if err := database.CreateUserAddress(&address); err != nil {
		return err
	}
	if err := database.UpdateUserStatus(id, "validated"); err != nil {
		return err
	}
	return nil
}

func UploadPhotoProfile(id uint, file *multipart.FileHeader) error {
	result, err := util.UploadFile(file)
	if err != nil {
		return err
	}
	if err := database.UploadPhotoProfil(id, result.Location); err != nil {
		return err
	}
	return nil
}
