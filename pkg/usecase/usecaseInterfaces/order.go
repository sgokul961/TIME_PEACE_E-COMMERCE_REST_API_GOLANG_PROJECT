package usecaseInterfaces

import "gokul.go/pkg/domain"

type OrderUseCase interface {
	GetOrders(id int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, addressid int, payment int) error
	CancelOrder(id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)
	EditOrderStatus(status string, id int) error
}
