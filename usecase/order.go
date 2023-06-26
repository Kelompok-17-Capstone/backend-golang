package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"net/http"

	"github.com/dariubs/percent"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

func CreateOrder(userId uint, req *payload.CreateOrder) error {
	var baseTotalPrice float64
	var grandTotalPrice float64
	var discount float64
	var coin int

	user, err := database.GetUser(userId)
	if err != nil {
		return err
	}

	for _, val := range req.Products {
		if val.Quantity <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Quantity must be greater than 0")
		}
	}

	if req.Products == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't order products, the products is nil")
	}

	for _, value := range req.Products {
		product, _ := database.GetProductById(value.ProductID)
		baseTotalPrice += float64(value.Quantity * int(product.Price))
	}

	grandTotalPrice = baseTotalPrice

	if user.Role == "member" {
		discount = percent.Percent(30, int(baseTotalPrice))
		grandTotalPrice = baseTotalPrice - discount
	}

	if req.Coin {
		if user.Coin > int(grandTotalPrice) {
			coin = int(grandTotalPrice)
		} else {
			coin = user.Coin
		}
		if err := database.CreateCoin(&models.Coin{
			UserID: userId,
			Total:  coin,
			Status: "decrease",
		}); err != nil {
			return err
		}
		grandTotalPrice = grandTotalPrice - float64(coin)

	}

	// validasi saldo dengan total price
	if user.Balance < int(grandTotalPrice) {
		return echo.NewHTTPError(http.StatusBadRequest, "the balance is not sufficient")
	}
	// update saldo dikurangi total price
	user.Balance -= int(grandTotalPrice)
	user.Coin -= coin

	if err := database.CreateTopup(&models.Balance{
		UserID: userId,
		Total:  int(grandTotalPrice),
		Status: "decrease",
	}); err != nil {
		return err
	}

	if err := database.UpdateUser(&user); err != nil {
		return err
	}
	order := models.Order{
		BaseTotalPrice:  int(baseTotalPrice),
		Discount:        int(discount),
		GrandTotalPrice: int(grandTotalPrice),
		Address:         req.Address,
		CoinsUsed:       coin,
		UserID:          userId,
		Status:          "dikemas",
	}
	orderId, err := database.CreateOrder(&order)
	if err != nil {
		return err
	}
	// update total product
	// create order dan order detail
	for _, value := range req.Products {
		product, err := database.GetProductById(value.ProductID)
		if err != nil {
			return err
		}
		product.Stock -= uint(value.Quantity)
		if err := database.UpdateProduct(value.ProductID, &product); err != nil {
			return err
		}
		orderDetail := models.OrderDetail{
			ProductID: value.ProductID,
			Quantity:  value.Quantity,
			OrderID:   orderId,
		}
		if err := database.CreateOrderDetail(&orderDetail); err != nil {
			return err
		}
		// hapus cart
		if req.Cart {
			if err := database.DeleteCartByProductId(value.ProductID); err != nil {
				return err
			}
		}
	}

	if err := database.SaveNotification(models.Notification{
		UserID: userId,
		Text:   "Pesanan anda telah berhasil dilakukan dan akan segera dikemas dengan id pesanan yaitu " + orderId.String(),
	}); err != nil {
		return err
	}

	return nil
}

func GetOrders(req *payload.OrdersParam) (resp []payload.GetOrders, err error) {
	orders, err := database.GetOrders(req)
	if err != nil {
		return
	}
	for _, value := range orders {
		var qty int
		var product string
		for key, val := range value.OrderDetails {
			qty += val.Quantity
			if key < len(value.OrderDetails)-1 {
				product += val.Product.Name + ", "
			} else {
				product += val.Product.Name
			}

		}
		resp = append(resp, payload.GetOrders{
			ID:            value.ID,
			Name:          value.User.Profile.Name,
			Address:       value.Address,
			Status:        value.Status,
			TotalQuantity: qty,
			TotalPrice:    value.GrandTotalPrice,
			OrderAt:       datatypes.Date(value.CreatedAt),
			ArrivedAt:     value.ArrivedAt,
			Products:      product,
		})
	}
	return
}
