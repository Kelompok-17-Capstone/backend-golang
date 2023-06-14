package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// update user email
func UpdateUserEmail(id uint, req *payload.UpdateEmail) error {
	if err := database.UpdateUserEmail(id, req.Email); err != nil {
		return err
	}
	return nil
}

// update user password
func UpdatePassword(id uint, req *payload.UpdatePassword) (models.User, error) {
	user, err := database.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return models.User{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid old password")
	}

	if req.NewPassword != req.RetypePassword {
		return models.User{}, echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(passwordHash)
	if err := database.Register(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateName(id uint, req *payload.Profile) error {
	if err := database.UpdateName(id, req.Name); err != nil {
		return err
	}
	return nil
}

// update phone number
func UpdatePhoneNumber(id uint, req *payload.UpdatePhoneNumber) error {
	if err := database.UpdatePhoneNumber(id, req.PhoneNumber); err != nil {
		return err
	}
	return nil
}

// update address
func UpdateAddress(id uint, addresId uint, req *payload.UpdateAddress) error {
	profile, err := database.GetProfile(id)
	if err != nil {
		return err
	}
	if req.Status == "primer" {
		if err := database.UpdateAddressStatus(profile.ID, addresId, "skunder"); err != nil {
			return err
		}
	}
	if err := database.UpdateAddress(profile.ID, addresId, req); err != nil {
		return err
	}
	return nil
}
