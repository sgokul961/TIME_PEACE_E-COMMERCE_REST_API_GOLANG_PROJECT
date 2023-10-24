package repository

import (
	"errors"
	"fmt"
	"strconv"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/utils/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminRepository{
		DB: DB,
	}

}
func (ad *AdminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select *from admins where email =? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}
	return adminCompareDetails, nil

}
func (ad *AdminRepository) GetUserById(id string) (domain.Users, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.Users{}, err
	}
	var count int

	if err := ad.DB.Raw("SELECT COUNT(*) FROM users WHERE id=?", user_id).Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}
	if count < 1 {
		return domain.Users{}, errors.New("user for the given id does not exist")
	}
	query := fmt.Sprintf("SELECT *FROM users WHERE id= '%d'", user_id)

	var userDeatils domain.Users

	if err := ad.DB.Raw(query).Scan(&userDeatils).Error; err != nil {
		return domain.Users{}, err
	}
	return userDeatils, nil
}

//function which will both block and unblock a user.

func (ad *AdminRepository) UpdateBlockUserID(user domain.Users) error {

	err := ad.DB.Exec("UPDATE users SET blocked = ? WHERE id = ?", user.Blocked, user.ID).Error

	if err != nil {
		fmt.Println("Error updating user :", err)
		return err
	}
	return nil

}
func (ad *AdminRepository) GetUsers(page int, count int) ([]models.UserDetailsAdmin, error) {
	//pagination pourpose

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count

	var userDeatils []models.UserDetailsAdmin

	if err := ad.DB.Raw("SELECT id, name, email,phone, blocked FROM users LIMIT ? OFFSET ?", count, offset).Scan(&userDeatils).Error; err != nil {
		return []models.UserDetailsAdmin{}, err
	}
	return userDeatils, nil
}

func (i *AdminRepository) Orderstatus(order_status string) ([]domain.Order, error) {

	var order_stat []domain.Order

	err := i.DB.Raw("SELECT *FROM orders WHERE order_status=?", order_status).Scan(&order_stat).Error
	if err != nil {
		return []domain.Order{}, err
	}
	return order_stat, nil
}
