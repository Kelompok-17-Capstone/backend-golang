package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"fmt"
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

func GetFavoriteProduct(userID uint) (*models.Product, error) {
	var product models.Product
	var user models.User

	err := config.DB.Preload("Favourites").Preload("Favourites.Product").First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	if len(user.Favorites) == 0 {
		return nil, fmt.Errorf("user has no favorite products")
	}

	product = user.Favorites[0].Product

	return &product, nil
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
