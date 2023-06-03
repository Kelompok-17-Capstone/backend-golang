package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func AddToCart(userID uint, productID uuid.UUID, quantity uint) error {
	product, err := database.GetProductById(productID)
	if err != nil {
		return fmt.Errorf("error getting product: %v", err)
	}
	if product.Stock < quantity {
		return fmt.Errorf("product stock is not available")
	}

	cart, err := database.GetCartItemByUserID(userID)
	if err != nil {
		cart = &models.CartItem{
			UserID: userID,
		}
		err = database.AddToCartItem(cart)
		if err != nil {
			return err
		}
	}
	var detailCart *models.DetailCartItem
	for _, val := range cart.DetailCartItems {
		if val.Products.ID == productID {
			detailCart = &val
			break
		}
	}

	if detailCart != nil {
		detailCart.Quantity += quantity
		err = database.UpdateDetailCartItem(detailCart.Quantity, detailCart.ID)
		if err != nil {
			return err
		}
	} else {
		detailCart = &models.DetailCartItem{
			CartItemID: cart.ID,
			ProductID:  product.ID,
			Quantity:   quantity,
		}
		err := database.AddToDetailCart(detailCart)
		if err != nil {
			return err
		}
	}
	// Cek apakah produk sudah ada di keranjang
	// var itemFound bool
	// for i, item := range cart.DetailCartItem {
	// 	if item.ProductID == productID {
	// 		cart.DetailCartItem[i].Quantity += quantity
	// 		itemFound = true
	// 		break
	// 	}
	// }

	// if !itemFound {
	// 	// Jika produk belum ada di keranjang, tambahkan item baru
	// 	item := models.DetailCartItem{
	// 		CartItemID: cart.ID,
	// 		ProductID:  productID,
	// 		Quantity:   quantity,
	// 	}
	// 	cart.DetailCartItem = append(cart.DetailCartItem, item)
	// 	err, _ := database.AddToDetailCart(detailCart.CartItemID)
	// 	if err != nil {
	// 		return fmt.Errorf("error adding")
	// 	}
	// 	if err := database.UpdateDetailCartItem(detailCart); err != nil {
	// 		return err
	// 	}
	// }
	// err = database.UpdateDetailCartItem(cart)
	// if err != nil {
	// 	return fmt.Errorf("error updating cart: %v", err)
	// }
	return nil
}
