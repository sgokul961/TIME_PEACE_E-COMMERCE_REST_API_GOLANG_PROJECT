package usecaseInterfaces

import (
	"gokul.go/pkg/domain"
	"gokul.go/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLoign) (models.TokenUsers, error)
	AddAddress(id int, address models.AddAddress) error
	GetAddress(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserDeatilsResponse, error)
	ChangePassword(id int, old string, password string, repassword string) error

	ForgotPasswordSend(phone string) error
	VarifyForgotPasswordAndChange(model models.ForgotVarify) error

	GetCart(id int) ([]models.GetCart, error)
	RemoveFromCart(cart_id, inventory_id int) error
	UpdateQuantityAdd(id, inventory_id int) error
	UpdateQuantityMinus(id, inventory_id int) error
}
