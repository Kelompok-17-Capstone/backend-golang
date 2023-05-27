package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"backend-golang/util"
	"mime/multipart"

	uuid "github.com/satori/go.uuid"
)

func UpdateProduct(id uuid.UUID, req *payload.UpdateProduct, image *multipart.FileHeader) (err error) {
	result, err := util.UploadFile(image)
	if err != nil {
		return err
	}
	if _, err := database.GetProductById(id); err != nil {
		return err
	}
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		Image:       result.Location,
	}
	if err := database.UpdateProduct(id, &product); err != nil {
		return err
	}
	return nil

}
