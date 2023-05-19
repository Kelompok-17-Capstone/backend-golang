package faker

import (
	"backend-golang/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("adminadmin"), bcrypt.DefaultCost)
	return &models.User{
		Email:    "admin@gmail.com",
		Password: string(password),
		Role:     "admin",
	}
}
