package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func SaveNotification(req models.Notification) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func CreateNotification(req *models.Notification) error {
	if err := config.DB.Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func GetNotificationById(id uint) (resp models.Notification, err error) {
	if err = config.DB.First(&resp, id).Error; err != nil {
		return
	}
	return
}

func GetNotifications(id uint) (resp []models.Notification, err error) {
	if err = config.DB.Order("created_at desc").Where("user_id = ?", id).Find(&resp).Error; err != nil {
		return
	}
	return
}

func DeleteNotificationById(id uint) error {
	if err := config.DB.Unscoped().Delete(&models.Notification{}, id).Error; err != nil {
		return err
	}
	return nil
}
