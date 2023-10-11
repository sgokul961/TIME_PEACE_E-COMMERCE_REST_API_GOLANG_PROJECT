package repository

import (
	"errors"
	"fmt"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/utils/models"
	"gorm.io/gorm"
)

type userDataBase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDataBase{DB: DB}
}
func (c *userDataBase) UserSignUp(user models.UserDetails) (models.UserDeatilsResponse, error) {
	var userDeatils models.UserDeatilsResponse

	err := c.DB.Raw("INSERT INTO users(name,email,password,phone)VALUES(?,?,?,?)RETURNING id,name,email,phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userDeatils).Error

	if err != nil {
		return models.UserDeatilsResponse{}, err
	}
	return userDeatils, nil
}
func (c *userDataBase) CheckUserAvailability(email string) bool {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*)FROM users WHERE EMAIL='%s'", email)

	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}
func (c *userDataBase) FindUserByEmail(user models.UserLoign) (models.UserSignInResponse, error) {

	var user_details models.UserSignInResponse

	err := c.DB.Raw(`SELECT * FROM users where email = ? and blocked =false`, user.Email).Scan(&user_details).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")

	}
	return user_details, nil
}
func (cr *userDataBase) UserBlockStatus(email string) (bool, error) {
	fmt.Println(email)

	var isBlocked bool

	err := cr.DB.Raw("SELECT blocked FROM users WHERE email = ?", email).Scan(&isBlocked).Error

	if err != nil {
		return false, err
	}
	fmt.Println(isBlocked)
	return isBlocked, nil
}
func (i *userDataBase) AddAddress(id int, address models.AddAddress) error {
	fmt.Println(id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin)

	err := i.DB.Exec(`INSERT INTO addresses(users_id ,name ,house_name,street,city,state,pin)
	VALUES($1, $2, $3, $4 ,$5, $6, $7)

	RETURNING id`, id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Error

	if err != nil {
		return err
	}
	return nil

}
func (ad *userDataBase) GetAddress(id int) ([]domain.Address, error) {
	var addresses []domain.Address

	if err := ad.DB.Where("users_id=?", id).Find(&addresses).Error; err != nil {
		return []domain.Address{}, err
	}
	return addresses, nil
}
func (ad *userDataBase) GetUserDetails(id int) (models.UserDeatilsResponse, error) {
	var details models.UserDeatilsResponse

	if err := ad.DB.Raw("SELECT id,name,email,phone FROM users WHERE id=?", id).Scan(&details).Error; err != nil {
		return models.UserDeatilsResponse{}, err
	}
	return details, nil
}
func (i *userDataBase) ChangePassword(id int, password string) error {
	err := i.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id).Error

	if err != nil {
		fmt.Println("error updating password", err)
		return err
	}
	return nil
}
func (i *userDataBase) GetPassword(id int) (string, error) {
	var userPassword string

	err := i.DB.Raw("SELECT password FROM users WHERE id=? ", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil
}
func (ad *userDataBase) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("select inventories.product_name,cart_products.quantity,cart_products.total_price from cart_products inner join inventories on cart_products.inventory_id=inventories.id where user_id=?", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}

	return cart, nil

}

func (ad *userDataBase) GetCartID(id int) (uint, error) {

	var cart_id uint
	if err := ad.DB.Raw("SELECT id FROM carts WHERE user_id=?", id).Scan(&cart_id).Error; err != nil {
		return 0, err
	}
	return cart_id, nil
}

// func (ad *userDataBase) GetProductsInCart(cart_id uint) ([]uint, error) {

// 	var cart_products []uint

// 	if err := ad.DB.Raw("SELECT inventory_id FROM line_items WHERE cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
// 		return nil, err
// 	}
// 	return cart_products, nil
// }

func (ad *userDataBase) GetProductsInCart(cart_id uint) ([]uint, error) {

	var cart_products []uint

	if err := ad.DB.Raw("SELECT inventory_id FROM line_items WHERE cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
		return nil, err
	}

	return cart_products, nil

}

func (ad *userDataBase) FindProductNames(inventory_id uint) (string, error) {

	var product_name string

	if err := ad.DB.Raw("SELECT product_name FROM inventories WHERE id =?", inventory_id).Scan(&product_name).Error; err != nil {
		return "", err

	}
	return product_name, nil

}
func (ad *userDataBase) FindPrice(inventory_id uint) (float64, error) {
	var price float64

	if err := ad.DB.Raw("SELECT price FROM inventories  WHERE id=?", inventory_id).Scan(&price).Error; err != nil {
		return 0, err
	}
	return price, nil
}
func (ad *userDataBase) FindCartQuantity(cart_id, inventory_id uint) (int, error) {
	var quantity int

	if err := ad.DB.Raw("SELECT quantity FROM line_items WHERE id=$1 AND inventory_id=$2", cart_id, inventory_id).Scan(&quantity).Error; err != nil {
		return 0, err
	}
	return quantity, nil
}

func (ad *userDataBase) FindCategory(inventory_id uint) (int, error) {
	var category int

	if err := ad.DB.Raw("SELECT category_id FROM inventories WHERE id=?", inventory_id).Scan(&category).Error; err != nil {
		return 0, err
	}
	return category, nil
}
func (ad *userDataBase) FindofferPercentage(category_id int) (int, error) {
	var percentage int
	err := ad.DB.Raw("select discount_rate from offers where category_id=$1 and valid=true", category_id).Scan(&percentage).Error
	if err != nil {
		return 0, err
	}

	return percentage, nil
}
func (ad *userDataBase) RemoveFromCart(cart_id, inventory_id int) error {

	if err := ad.DB.Exec("DELETE FROM line_items WHERE cart_id=$1 AND inventory_id= $2", cart_id, inventory_id).Error; err != nil {
		return err
	}
	return nil
}
func (ad *userDataBase) UpdateQuantityAdd(id, inventory_id int) error {

	query := `UPDATE line_items SET quantity = quantity +1
	WHERE cart_id=$1 AND inventory_id=$2`
	result := ad.DB.Exec(query, id, inventory_id)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (ad *userDataBase) UpdateQuantityMinus(id, inventory_id int) error {
	query := `UPDATE line_items SET quantity =quantity-1
WHERE cart_id=$1 AND inventory_id=$2`
	result := ad.DB.Exec(query, id, inventory_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (ad *userDataBase) FindIdFromPhone(phone string) (int, error) {

	var id int

	if err := ad.DB.Raw(`SELECT id FROM users WHERE phone=?`, phone).Scan(&id).Error; err != nil {
		return id, err

	}
	return id, nil
}
