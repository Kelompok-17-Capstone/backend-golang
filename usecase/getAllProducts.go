package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"

	uuid "github.com/satori/go.uuid"
)

func GetAllProduct(status string, keyword string) (resp []payload.ProductResponse, err error) {
	var products []models.Product
	if keyword == "" {
		if status == "" {
			p, e := database.GetAllProduct()
			products = p
			err = e
		} else if status == "tersedia" {
			p, e := database.GetProductsByStock(">")
			products = p
			err = e
		} else if status == "habis" {
			p, e := database.GetProductsByStock("=")
			products = p
			err = e
		}
	} else {
		if status == "" {
			p, e := database.GetProductsByName(keyword)
			products = p
			err = e
		} else if status == "tersedia" {
			p, e := database.GetProductsByNameAndStock(keyword, ">")
			products = p
			err = e
		} else if status == "habis" {
			p, e := database.GetProductsByNameAndStock(keyword, "=")
			products = p
			err = e
		}
	}

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
