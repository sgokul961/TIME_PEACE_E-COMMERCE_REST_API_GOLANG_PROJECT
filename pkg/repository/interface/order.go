package interfaces

import (
	"gokul.go/pkg/domain"
	"gokul.go/pkg/utils/models"
)

type OrderRepository interface {
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	GetOrders(id int) ([]domain.Order, error)
	//GetCart(userid int) ([]models.GetCart,error)

	AdminOrders(status string) ([]domain.OrderDetails, error)
	EditOrderStatus(status string, id int) error

	AddOrderProducts(order_id int, cart []models.GetCart) error
	CancelOrder(id int) error
}
