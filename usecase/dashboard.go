package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"

	"gorm.io/datatypes"
)

func GetDashboard() (resp []payload.DashboardResponse, err error) {
	users, err := database.DashboardUsers()
	if err != nil {
		return nil, err
	}

	orders, err := database.DashboardOrders()
	if err != nil {
		return nil, err
	}

	products, err := database.DashboardProducts()
	if err != nil {
		return nil, err
	}

	orderDetails, err := database.DashboardOrderDetails()
	if err != nil {
		return nil, err
	}

	var usersResponse []payload.UsersResponse
	for _, user := range users {
		var address string
		if len(user.Profile.Address) > 0 {
			address = user.Profile.Address[0].Address + ", " + user.Profile.Address[0].City + ", " + user.Profile.Address[0].Province
		}
		usersResponse = append(usersResponse, payload.UsersResponse{
			Name:        user.Profile.Name,
			PhoneNumber: user.Profile.PhoneNumber,
			Address:     address,
			Role:        user.Role,
		})
	}

	var ordersResponse []payload.OrdersResponse
	for _, order := range orders {
		var productName string
		var quantity int
		for key, orderDetail := range order.OrderDetails {
			if key < len(order.OrderDetails)-1 {
				productName += orderDetail.Product.Name + ", "
			} else {
				productName += orderDetail.Product.Name
			}
			quantity += orderDetail.Quantity
		}
		ordersResponse = append(ordersResponse, payload.OrdersResponse{
			UserName:    order.User.Profile.Name,
			Address:     order.Address,
			ProductName: productName,
			Quantity:    quantity,
			OrderAt:     datatypes.Date(order.CreatedAt),
		})
	}

	var productsResponse []payload.ProductsResponse
	for _, product := range products {
		productsResponse = append(productsResponse, payload.ProductsResponse{
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			Status:      product.Status,
		})
	}

	resp = append(resp, payload.DashboardResponse{
		Users:    usersResponse,
		Orders:   ordersResponse,
		Products: productsResponse,
		Graphic:  orderDetails,
	})

	return resp, nil
}
