package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"
)

// update user email
func UpdateUserEmail(id uint, email string) error {
	var user models.User
	if err := config.DB.Model(&user).Where("id = ?", id).Update("email", email).Error; err != nil {
		return err
	}
	return nil
}

// update user password
func GetUserById(id uint) (user models.User, err error) {
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return
}

// upadate Name User
func UpdateName(id uint, name string) error {
	if err := config.DB.Model(&models.Profile{}).Where("user_id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

// update phone number
func UpdatePhoneNumber(id uint, phoneNumber string) error {
	if err := config.DB.Model(&models.Profile{}).Where("user_id = ?", id).Update("phone_number", phoneNumber).Error; err != nil {
		return err
	}
	return nil
}

// update address
func UpdateAddress(id uint, addresId uint, req *payload.UpdateAddress) error {
	if err := config.DB.Model(&models.Address{}).Where("profile_id = ? and id = ? ", id, addresId).Updates(models.Address{
		Address:  req.Address,
		City:     req.City,
		Province: req.Province,
		Status:   req.Status,
	}).Error; err != nil {
		return err
	}
	return nil
}

func SaveAddres(req *models.Address) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAddressWhereStatus(profilId uint, req *models.Address, status bool) error {
	if err := config.DB.Model(&models.Address{}).Where("profile_id = ? and status = ? ", profilId, status).Updates(models.Address{
		Address:  req.Address,
		City:     req.City,
		Province: req.Province,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GetAddressById(id uint) (resp models.Address, err error) {
	if err = config.DB.First(&resp, id).Error; err != nil {
		return
	}
	return
}

// get profile by id
func GetProfile(id uint) (resp models.Profile, err error) {
	if err := config.DB.Where("user_id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}
