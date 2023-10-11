package usecaseInterfaces

import "gokul.go/pkg/utils/models"

type CartUseCase interface {
	AddToCart(user_id, inventory_id int) error
	CheckOut(id int) (models.CheckOut, error)
}
