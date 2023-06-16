package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"time"

	uuid "github.com/satori/go.uuid"
)

// update status order
func UpdateStatusOrder(id uuid.UUID, req *payload.UpdateOrderStatus) error {
	if err := database.UpdateStatusOrder(id, req.Status); err != nil {
		return err
	}
	return nil
}

// update arrived at order
func UpdateArrivedAtOrder(id uuid.UUID, req *payload.UpadateArrivedAt) error {
	arrivedAt, err := time.Parse(time.RFC3339, req.ArrivedAt)
	if err != nil {
		return err
	}

	if err := database.UpdateArrivedAtOrder(id, arrivedAt); err != nil {
		return err
	}
	return nil

}
