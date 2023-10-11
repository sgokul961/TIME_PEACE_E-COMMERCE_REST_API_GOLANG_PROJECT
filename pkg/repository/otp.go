package repository

import (
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/utils/models"
	"gorm.io/gorm"
)

type OtpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &OtpRepository{
		DB: DB,
	}
}
func (ot *OtpRepository) FindUserByMobileNumber(phone string) bool {
	var count int

	if err := ot.DB.Raw("SELECT COUNT(*) FROM users WHERE phone = ?", phone).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}
func (ot *OtpRepository) UserDeatilsUsingPhone(phone string) (models.UserDeatilsResponse, error) {
	var userDetails models.UserDeatilsResponse

	if err := ot.DB.Raw("SELECT *FROM users WHERE phone =?", phone).Scan(&userDetails).Error; err != nil {
		return models.UserDeatilsResponse{}, err
	}
	return userDetails, nil
}
