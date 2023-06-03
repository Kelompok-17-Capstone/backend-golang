package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"

	uuid "github.com/satori/go.uuid"
)

// get all product

func GetProducts(req *payload.ProductParam) ([]models.Product, error) {
	var products []models.Product
	db := config.DB
	if req.Keyword != "" {
		db = db.Where("name like ?", "%"+req.Keyword+"%")
	}
	if req.Status != "" {
		if req.Status == "tersedia" {
			db = db.Where("stock > 0")
		} else if req.Status == "habis" {
			db = db.Where("stock = 0")
		}
	}
	if err := db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

// get product by id

func GetProductById(id uuid.UUID) (resp models.Product, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetProductsMobile(req *payload.ProductParam) ([]models.Product, error) {
	var products []models.Product
	db := config.DB
	if req.Keyword != "" {
		db = db.Where("name like ?", "%"+req.Keyword+"%")
	}
	if req.Tab != "" {
		if req.Tab == "terbaru" {
			db = db.Order("created_at desc")
		} else if req.Tab == "terfavorit" {
			db = db.Order("favorite desc")
		}
	}
	if req.Price != "" {
		db = db.Order("price " + req.Price)
	}
	if err := db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}
