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
	favoriteProduct, err := database.GetFavoriteProduct(userID)
	if err != nil {
		return nil, err
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
