package usecaseInterfaces

import (
	"github.com/jung-kurt/gofpdf"
	"gokul.go/pkg/domain"
)

type OrderUseCase interface {
	GetOrders(id int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, addressid int, payment int, couponID int) error
	CancelOrder(id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)
	EditOrderStatus(status string, id int) error
	GenerateInvoice(orderID uint) (*gofpdf.Fpdf, error)
}
