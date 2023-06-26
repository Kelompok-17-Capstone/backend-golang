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
		Status:    true,
	}
	if err := database.CreateUserAddress(&address); err != nil {
		return err
	}
	if err := database.UpdateUserStatus(id, "validated"); err != nil {
		return err
	}
	user, err := database.GetUser(id)
	if err != nil {
		return err
	}
	user.Balance += 10000
	if err := database.UpdateUser(&user); err != nil {
		return err
	}
	if err := database.SaveNotification(models.Notification{
		UserID: id,
		Text:   "Selamat! Anda mendapatkan bonus saldo sebesar Rp10.000 karena telah mendaftarkan akun di aplikasi mobile AltaTech",
	}); err != nil {
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
