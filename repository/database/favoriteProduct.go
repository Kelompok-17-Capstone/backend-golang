package database

import (
	"backend-golang/config"
	"backend-golang/models"

	uuid "github.com/satori/go.uuid"
)

func CheckFavoritProduct(productId uuid.UUID, userId uint) error {
	if err := config.DB.Where("product_id = ? and user_id = ?", productId, userId).First(&models.FavoriteProduct{}).Error; err != nil {
		return err
	}
	return nil
}

func CheckFavoriteIdAndUserId(userID uint, id uint) error {
	if err := config.DB.Where("user_id =  ?", userID).First(&models.FavoriteProduct{}, id).Error; err != nil {
		return err
	}

	return nil
}

func AddFavoriteProduct(fav *models.FavoriteProduct) error {
	if err := config.DB.Save(&fav).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFavoriteProduct(userID uint, id uint) error {
	var fav models.FavoriteProduct
	config.DB.First(&fav, id)
	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&models.FavoriteProduct{ProductID: fav.ProductID}).Error; err != nil {
		return err
	}
	return nil
}
