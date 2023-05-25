package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

func CreateProduct(req *payload.CreateProduct, image *multipart.FileHeader) error {
	product := models.Product{
		ProductName: req.ProductName,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
	}
	if image != nil {
		fileName := uuid.NewV4().String() + filepath.Ext(image.Filename)
		uploadPath := filepath.Join("images", fileName)
		file, err := image.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		dst, err := os.Create(uploadPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, file); err != nil {
			return err
		}

		product.ImageFile = uploadPath
	}
	if err := database.CreateProduct(&product); err != nil {
		return err
	}
	return nil
}
