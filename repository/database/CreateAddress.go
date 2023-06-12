package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

// create address
// func CreateAddress(profileID uint, req *payload.CreateAddress) error {
// 	address := models.Address{
// 		ProfileID: profileID,
// 		Address:   req.Address,
// 		City:      req.City,
// 		Province:  req.Province,
// 	}
// 	if err := config.DB.Create(&address).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func CreateAddress(req *models.Address) error {
	if err := config.DB.Create(&req).Error; err != nil {
		return err
	}
	return nil
}
