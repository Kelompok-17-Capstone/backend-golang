package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"time"

	uuid "github.com/satori/go.uuid"
)

// delete order by id
func DeleteOrder(id uuid.UUID) error {
	var order models.Order
	if err := config.DB.Model(&order).Where("id = ?", id).Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

// update order by ID
func UpdateOrderStatusAndArrived(id uuid.UUID, status string, arrivedAt time.Time) error {
	if err := config.DB.Model(&models.Order{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":     status,
		"arrived_at": arrivedAt,
	}).Error; err != nil {
		return err
	}
	return nil
}

// get order by id
func GetOrderByID(id uuid.UUID) (*models.Order, error) {
	var order models.Order
	if err := config.DB.Where("id = ?", id).Preload("OrderDetails.Product").First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// detele order detail by order id
func DeleteOrderDetailByOrderID(orderID uuid.UUID) error {
	if err := config.DB.Where("order_id = ?", orderID).Delete(&models.OrderDetail{}).Error; err != nil {
		return err
	}
	return nil
}
