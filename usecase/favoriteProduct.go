package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func AddFavoriteProduct(userID uint, productID string) error {
	err := database.AddFavoriteProduct(userID, productID)
	if err != nil {
		return err
	}

	return nil
}

func GetFavoriteProduct(userID uint) (*payload.GetFavoriteProduct, error) {
	product, err := database.GetFavoriteProduct(userID)
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
	err := database.DeleteFavoriteProduct(userID, productID)
	if err != nil {
		return err
	}

	return nil
}
