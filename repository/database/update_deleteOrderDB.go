package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"time"

	uuid "github.com/satori/go.uuid"
)

// update status order
func UpdateStatusOrder(id uuid.UUID, status string) error {
	if err := config.DB.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

// update arrived at order dengan parsing string ke time
func UpdateArrivedAtOrder(id uuid.UUID, arrivedAt time.Time) error {
	if err := config.DB.Model(&models.Order{}).Where("id = ?", id).Update("arrived_at", arrivedAt).Error; err != nil {
		return err
	}
	return nil
}
