package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func AddFavoriteProduct(userID uint, productID string) error {
	productUUID := uuid.FromStringOrNil(productID)
	if err := database.CheckFavoritProduct(productUUID, userID); err == nil {
		return fmt.Errorf("cant favourite same product")
	}
	fav := models.FavoriteProduct{
		UserID:    userID,
		ProductID: productUUID,
	}
	err := database.AddFavoriteProduct(&fav)
	if err != nil {
		return err
	}

	return nil
}

func GetFavoriteProduct(userID uint) ([]payload.GetFavoriteProduct, error) {
	var resp []payload.GetFavoriteProduct
	user, err := database.GetUser(userID)
	if err != nil {
		return nil, err
	}

	if len(user.Favorites) == 0 {
		return nil, fmt.Errorf("user has no favorite products")
	}

	for _, value := range user.Favorites {
		resp = append(resp, payload.GetFavoriteProduct{
			ID:          value.ID,
			ProductID:   value.Product.ID,
			Name:        value.Product.Name,
			Description: value.Product.Description,
			Price:       value.Product.Price,
			Image:       value.Product.Image,
			Favorite:    true,
		})
	}

	return resp, nil
}

func DeleteFavoriteProduct(userID uint, id uint) error {
	if err := database.CheckFavoriteIdAndUserId(userID, id); err != nil {
		return fmt.Errorf("failed to delete favorite product")
	}

	if err := database.DeleteFavoriteProduct(userID, id); err != nil {
		return err
	}

	return nil
}
