package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"time"

	uuid "github.com/satori/go.uuid"
)

// delete order by id
func DeleteOrder(id uuid.UUID) error {
	if err := database.DeleteOrder(id); err != nil {
		return err
	}
	return nil
}

// update order by ID
func UpdateOrderStatusAndArrived(id uuid.UUID, req *payload.UpdateOrder) error {
	arrivedAt, err := time.Parse(time.RFC3339, req.ArrivedAt)
	if err != nil {
		return err
	}

	if err := database.UpdateOrderStatusAndArrived(id, req.Status, arrivedAt); err != nil {
		return err
	}
	return nil
}

// get order by id
func GetOrderByID(id uuid.UUID) (*payload.GetOrders, error) {
	order, err := database.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	var qty int
	var product []string
	for _, val := range order.OrderDetails {
		qty += val.Quantity
		product = append(product, val.Product.Name)

	}

	respon := &payload.GetOrders{
		Name:          order.User.Profile.Name,
		Address:       order.Address,
		TotalQuantity: qty,
		TotalPrice:    order.GrandTotalPrice,
		OrderAt:       order.OrderAt,
		ArrivedAt:     order.ArrivedAt,
		Status:        order.Status,
	}

	return respon, nil
}
