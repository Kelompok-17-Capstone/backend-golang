package database

import (
	"backend-golang/config"
	"backend-golang/models"

	"fmt"

	uuid "github.com/satori/go.uuid"
)

func ViewMemberInformation(id uint) (models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).Preload("Profile.Address").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GenerateMemberCode(id uint) string {
	uuidObj := uuid.NewV5(uuid.NamespaceOID, fmt.Sprintf("%d", id))
	return uuidObj.String()
}

func StoreMemberCode(id uint, memberCode string) error {
	var user models.User
	if err := config.DB.Model(&user).Where("id = ?", id).Update("member_code", memberCode).Error; err != nil {
		return err
	}
	return nil
}

func GetMemberCode(id uint) (string, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return "", err
	}
	return user.MemberCode, nil
}
