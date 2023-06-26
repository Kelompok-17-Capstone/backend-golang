package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

// delete order by id
func DeleteOrder(id uuid.UUID) error {
	if _, err := database.GetOrderByID(id); err != nil {
		return err
	}
	if err := database.DeleteOrder(id); err != nil {
		return err
	}
	if err := database.DeleteOrderDetailByOrderID(id); err != nil {
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
	order, err := database.GetOrderByID(id)
	if err != nil {
		return err
	}
	if err := database.UpdateOrderStatusAndArrived(id, req.Status, arrivedAt); err != nil {
		return err
	}

	if req.Status != "" {
		if err := database.SaveNotification(models.Notification{
			UserID: order.UserID,
			Text:   "Pesanan anda dengan id pesanan " + id.String() + " telah " + req.Status,
		}); err != nil {
			return err
		}
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
	var product string
	for key, val := range order.OrderDetails {
		qty += val.Quantity
		if key < len(order.OrderDetails)-1 {
			product += val.Product.Name + ", "
		} else {
			product += val.Product.Name
		}

	}

	respon := &payload.GetOrders{
		ID:            order.ID,
		Name:          order.User.Profile.Name,
		Address:       order.Address,
		TotalQuantity: qty,
		TotalPrice:    order.GrandTotalPrice,
		Products:      product,
		OrderAt:       datatypes.Date(order.CreatedAt),
		ArrivedAt:     order.ArrivedAt,
		Status:        order.Status,
	}

	return respon, nil
}
