package interfaces

import "gokul.go/pkg/utils/models"

type CartRepository interface {
	GetCart(id int) ([]models.GetCart, error)
	GetCartId(cart_id int) (int, error)
	CreateNewCart(cart_id int) (int, error)
	AddLineItems(cart_id, inventory_id int) error
	GetAddresses(id int) ([]models.Address, error)
	GetPaymentOptions() ([]models.PaymentMethod, error)

	//-------------------------------------------------------//
}
