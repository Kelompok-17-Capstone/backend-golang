package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

// import (
// 	"backend-golang/models"
// 	"backend-golang/repository/database"
// )

// func GetAllProduct() ([]models.Product, error) {
// 	products, err := database.GetAllProduct()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return products, nil
// }

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
