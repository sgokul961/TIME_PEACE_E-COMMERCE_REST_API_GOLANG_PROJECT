package repository

import (
	interfaces "gokul.go/pkg/repository/interface"
	"gorm.io/gorm"
)

type paymentRepositoy struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.PayementRepository {
	return &paymentRepositoy{
		DB: db,
	}
}

func (p *paymentRepositoy) FindUserName(user_id int) (string, error) {
	var name string

	if err := p.DB.Raw(`SELECT name FROM users WHERE id=?`, user_id).Scan(&name).Error; err != nil {
		return "", err
	}
	return name, nil
}
func (p *paymentRepositoy) FindPrice(order_id int) (float64, error) {
	var price float64

	if err := p.DB.Raw("SELECT final_price FROM orders WHERE id=?", order_id).Scan(&price).Error; err != nil {
		return 0, err
	}
	return price, nil
}
func (p *paymentRepositoy) UpdatePaymentDetails(OrderID, paymentID, razorID string) error {
	status := "PAID"

	if err := p.DB.Exec(`UPDATE orders SET payment_status =$1 WHERE id =$2`, status, OrderID).Error; err != nil {
		return err
	}
	return nil
}
