package interfaces

import (
	"gokul.go/pkg/domain"
	"gokul.go/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDeatilsResponse, error)
	CheckUserAvailability(email string) bool
	FindUserByEmail(user models.UserLoign) (models.UserSignInResponse, error)
	UserBlockStatus(email string) (bool, error)
	AddAddress(id int, address models.AddAddress) error
	GetAddress(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserDeatilsResponse, error)
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
	FindIdFromPhone(phone string) (int, error)

	GetCart(id int) ([]models.GetCart, error)
	RemoveFromCart(cart_id, inventory_id int) error
	UpdateQuantityAdd(id, inventory_id int) error
	UpdateQuantityMinus(id, inventory_id int) error

	GetCartID(id int) (uint, error)
	GetProductsInCart(cart_id uint) ([]uint, error)
	FindProductNames(inventory_id uint) (string, error)
	FindPrice(inventory_id uint) (float64, error)
	FindCartQuantity(cart_id, inventory_id uint) (int, error)
	FindCategory(inventory_id uint) (int, error)
	FindofferPercentage(category_id int) (int, error)
}
