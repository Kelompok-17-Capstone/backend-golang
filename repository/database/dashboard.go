package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"
)

func DashboardUsers() ([]*models.User, error) {
	var user []*models.User
	if err := config.DB.Preload("Profile.Address").Not("role = ? ", "admin").Order("created_at desc").Find(&user).Limit(4).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DashboardOrders() ([]*models.Order, error) {
	var order []*models.Order
	if err := config.DB.Preload("User.Profile.Address").Preload("OrderDetails.Product").Order("created_at desc").Find(&order).Limit(4).Error; err != nil {
		return order, err
	}
	return order, nil
}

func DashboardProducts() ([]*models.Product, error) {
	var product []*models.Product
	if err := config.DB.Order("created_at desc").Find(&product).Limit(4).Error; err != nil {
		return product, err
	}
	return product, nil
}

func DashboardOrderDetails() (resp []payload.Graphic, err error) {
	config.DB.Table("order_details").
		Select("products.name, SUM(order_details.quantity) as qty").
		Joins("JOIN products ON products.id = order_details.product_id").Group("products.name").
		Order("qty DESC").Limit(6).
		Scan(&resp)
	return
}
