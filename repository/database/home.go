package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetLimitsProductsOrderByCreatedAt(limit int, order string) (resp []models.Product, err error) {
	if err = config.DB.Limit(limit).Order("created_at " + order).Find(&resp).Error; err != nil {
		return
	}
	return
}
