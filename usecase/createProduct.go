package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"backend-golang/util"
	"mime/multipart"
)

func CreateProduct(req *payload.CreateProduct, image *multipart.FileHeader) error {
	result, err := util.UploadFile(image)
	if err != nil {
		return err
	}
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		Image:       result.Location,
	}
	if err := database.CreateProduct(&product); err != nil {
		return err
	}
	return nil
}
