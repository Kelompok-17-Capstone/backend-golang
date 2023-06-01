package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"
)

func AddFavoriteProduct(userID uint, productID string) error {
	var product models.FavoriteProduct
	err := config.DB.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return err
	}

	product.UserID = userID

	err = config.DB.Save(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func GetFavoriteProduct(userID uint) (*payload.GetFavoriteProduct, error) {
	var product models.Product
	err := config.DB.Where("user_id = ? AND favorite = ?", userID, true).First(&product).Error
	if err != nil {
		return nil, err
	}

	favoriteProduct := &payload.GetFavoriteProduct{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		Favorite:    true,
	}

	return favoriteProduct, nil
}

func DeleteFavoriteProduct(userID uint, productID string) error {
	var product models.FavoriteProduct
	err := config.DB.Where("user_id = ? AND id = ?", userID, productID).First(&product).Error
	if err != nil {
		return err
	}

	err = config.DB.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}
