package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func Home(userId uint) (resp payload.GetHome, err error) {
	user, err := database.GetUser(userId)
	if err != nil {
		return
	}
	product, err := database.GetLimitsProductsOrderByCreatedAt(6, "desc")

	if err != nil {
		return
	}
	var getProduct []payload.ProductsHome

	for _, value := range product {
		getProduct = append(getProduct, payload.ProductsHome{
			Name:  value.Name,
			Image: value.Image,
		})
	}

	resp = payload.GetHome{
		Saldo:    user.Balance,
		Products: getProduct,
	}
	return
}

func Guest() (resp payload.GetHome, err error) {
	product, err := database.GetLimitsProductsOrderByCreatedAt(6, "desc")

	if err != nil {
		return
	}
	var getProduct []payload.ProductsHome

	for _, value := range product {
		getProduct = append(getProduct, payload.ProductsHome{
			Name:  value.Name,
			Image: value.Image,
		})
	}

	resp = payload.GetHome{
		Saldo:    0,
		Products: getProduct,
	}
	return
}
