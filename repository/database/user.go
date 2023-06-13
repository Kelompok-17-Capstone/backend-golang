package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetUserByEmail(email string) (user models.User, err error) {
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return
}

func GetUserProfile(id uint) (resp *models.Profile, err error) {
	if err := config.DB.Where("user_id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func CreateUserProfil(req *models.Profile) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserProfil(req *models.Profile, id uint) error {
	if err := config.DB.Model(&req).Where("user_id = ?", id).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

// func GetAdress(profilId uint) error {
// 	if err := DB.Where("profil_id = ?",profilId).First(&req,id).Error
// }

func CreateUserAddress(req *models.Address) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserAddress(req *models.Address, id uint) error {
	if err := config.DB.Model(&req).Where("profile_id = ?", id).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAddressStatus(profileId uint, addressId uint, status bool) error {
	if err := config.DB.Model(&models.Address{}).Where("profile_id = ?", profileId).Not("id = ?", addressId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserStatus(id uint, status string) error {
	var user models.User
	if err := config.DB.Model(&user).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func UploadPhotoProfil(id uint, photo string) error {
	var profile *models.Profile
	if err := config.DB.Model(&profile).Where("user_id = ?", id).Update("photo", photo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAddress(id uint) error {
	if err := config.DB.Delete(&models.Address{}, id).Error; err != nil {
		return err
	}
	return nil
}
