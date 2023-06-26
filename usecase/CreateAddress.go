package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

// CreateAddress creates a new address for a profile
// func CreateAddress(userID uint, req *payload.CreateAddress) error {
// 	// Get the profile by ID
// 	profile, err := database.GetProfile(userID)
// 	if err != nil {
// 		return err
// 	}

// 	// Create the address
// 	err = database.CreateAddress(profile.ID, req)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func CreateAddress(profielID uint, req *payload.CreateAddress) error {
// 	address := models.Address{
// 		ProfileID: profielID,
// 		Address:   req.Address,
// 		City:      req.City,
// 		Province:  req.Province,
// 	}
// 	if err := database.CreateAddress(&address); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func CreateAddress(profileID uint, req *payload.CreateAddress) error {
// 	// Get the profile by ID
// 	profile, err := database.GetProfile(profileID)
// 	if err != nil {
// 		return err
// 	}

// 	address := models.Address{
// 		ProfileID: profile.ID,
// 		Address:   req.Address,
// 		City:      req.City,
// 		Province:  req.Province,
// 	}
// 	if err := database.CreateAddress(&address); err != nil {
// 		return err
// 	}
// 	return nil
// }

func CreateAddress(id uint, req *payload.CreateAddress) error {
	// Get the profile by ID
	profile, err := database.GetProfile(id)
	if err != nil {
		return err
	}

	address := models.Address{
		ProfileID: profile.ID,
		Address:   req.Address,
		City:      req.City,
		Province:  req.Province,
		Status:    false,
	}
	if err := database.CreateAddress(&address); err != nil {
		return err
	}
	return nil
}
