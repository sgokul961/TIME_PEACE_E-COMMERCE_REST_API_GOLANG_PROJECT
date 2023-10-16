package repository

import (
	"time"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gorm.io/gorm"
)

type SalesRepository struct {
	DB *gorm.DB
}

func NewSalesRepository(db *gorm.DB) interfaces.SalesRepository {
	return &SalesRepository{DB: db}
}
func (s *SalesRepository) GetMonthlySalesReport(year int, month time.Month) ([]domain.Order, error) {
	var orders []domain.Order

	result := s.DB.
		Preload("OrderItems").
		Preload("PaymentMethod").
		Preload("OrderItems.Inventories").
		Preload("Users").
		Where("EXTRACT(YEAR FROM created_at)=? AND EXTRACT(MONTH FROM created_at)=?", year, month).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil

}
