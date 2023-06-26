package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"
)

func GetOrderMobile(id uint, req *payload.GetOrderMobileParam) ([]models.Order, error) {
	var orders []models.Order
	db := config.DB
	if req.Status != "" {
		if req.Status == "dikemas" || req.Status == "dikirim" || req.Status == "diterima" {
			db = db.Where("status = ?", req.Status).Order("created_at DESC")
		}
	}

	if err := db.Preload("OrderDetails.Product").Where("user_id = ?", id).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}
