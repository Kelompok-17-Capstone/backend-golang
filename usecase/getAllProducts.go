package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"

	uuid "github.com/satori/go.uuid"
)

// get all product

func GetAllProduct() (resp []payload.ProductResponse, err error) {
	products, err := database.GetAllProduct()
	if err != nil {
		return []payload.ProductResponse{}, err
	}
	for _, product := range products {
		resp = append(resp, payload.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			Image:       product.Image,
		})
	}
	return
}

// get prodcut by ID

func GetProductByid(id uuid.UUID) (resp payload.ProductResponse, err error) {
	product, err := database.GetProductById(id)
	if err != nil {
		return payload.ProductResponse{}, err
	}
	resp = payload.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		Image:       product.Image,
	}
	return
}
