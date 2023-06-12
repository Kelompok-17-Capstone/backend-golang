package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"

	uuid "github.com/satori/go.uuid"
)

func GetProducts(status string, keyword string) (resp []payload.ProductResponse, err error) {
	req := payload.ProductParam{
		Keyword: keyword,
		Status:  status,
	}
	products, err := database.GetProducts(&req)
	if err != nil {
		return []payload.ProductResponse{}, err
	}
	for _, product := range products {
		var status string
		if product.Stock > 0 {
			status = "tersedia"
		} else {
			status = "habis"
		}
		resp = append(resp, payload.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			Image:       product.Image,
			Status:      status,
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

func GetProductsMobile(keyword, tab, price string) (resp []payload.ProductMobileResponse, err error) {
	req := payload.ProductParam{
		Keyword: keyword,
		Tab:     tab,
		Price:   price,
	}
	products, err := database.GetProductsMobile(&req)
	if err != nil {
		return resp, err
	}
	for _, product := range products {
		var status string
		if product.Stock > 0 {
			status = "tersedia"
		} else {
			status = "habis"
		}
		resp = append(resp, payload.ProductMobileResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			Image:       product.Image,
			Status:      status,
			Favorit:     int(product.Favorite),
		})
	}
	return
}

func GetProductMobileByid(id uuid.UUID) (resp payload.ProductMobileResponse, err error) {
	product, err := database.GetProductById(id)
	if err != nil {
		return
	}
	resp = payload.ProductMobileResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		Image:       product.Image,
		Favorit:     int(product.Favorite),
	}
	return
}
