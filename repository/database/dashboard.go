package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"
)

func DashboardUsers() ([]*models.User, error) {
	var user []*models.User
	if err := config.DB.Preload("Profile.Address").Not("role = ? ", "admin").Find(&user).Limit(4).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DashboardOrders() ([]*models.Order, error) {
	var order []*models.Order
	if err := config.DB.Preload("User.Profile.Address").Preload("OrderDetails.Product").Find(&order).Limit(4).Error; err != nil {
		return order, err
	}
	return order, nil
}

func DashboardProducts() ([]*models.Product, error) {
	var product []*models.Product
	if err := config.DB.Find(&product).Limit(4).Error; err != nil {
		return product, err
	}
	return product, nil
}

func DashboardOrderDetails() (string, int, error) {
	var orderDetail []*models.OrderDetail
	var graphic []*payload.Graphic
	if err := config.DB.Model(orderDetail).
		Select("products.name, SUM(order_details.quantity) as qty").
		Joins("JOIN products ON products.id = order_details.product_id").
		Order("qty DESC").
		Group("products.name").Row().Scan().Error; err != nil {
		return graphic, err
	}
	return graphic, nil
}
