package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
	"backend-golang/util"
	"mime/multipart"

	uuid "github.com/satori/go.uuid"
)

func UpdateProduct(id uuid.UUID, image *multipart.FileHeader) (product *models.Product, err error) {
	result, err := util.UploadFile(image)
	if err != nil {
		return nil, err
	}
	product = &models.Product{
		Image: result.Location,
	}
	err = database.UpdateProduct(id, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
