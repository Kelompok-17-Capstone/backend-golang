package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"

	uuid "github.com/satori/go.uuid"
)

func CreateOrder(order *models.Order) (orderId uuid.UUID, err error) {
	if err = config.DB.Save(&order).Error; err != nil {
		return
	}
	return order.ID, nil
}

func CreateOrderDetail(req *models.OrderDetail) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func GetOrders(req *payload.OrdersParam) (orders []models.Order, err error) {
	db := config.DB
	if req.Keyword != "" {
		db = db.Joins("JOIN orders ON orders.user_id = users.id").Where("users.name like ?", "%"+req.Keyword+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if err = db.Preload("User.Profile.Address").Preload("OrderDetails.Product").Find(&orders).Error; err != nil {
		return
	}
	return
}
