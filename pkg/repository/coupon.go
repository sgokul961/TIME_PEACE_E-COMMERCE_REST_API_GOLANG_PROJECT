package repository

import (
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/utils/models"
	"gorm.io/gorm"
)

type CouponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.CouponRepository {
	return &CouponRepository{
		DB: db,
	}
}
func (repo *CouponRepository) Addcoupon(coup models.Coupons) error {
	if err := repo.DB.Exec(`INSERT INTO coupons(coupon,discount_rate,valid)values($1,$2,$3)`, coup.Coupon, coup.DiscountRate, coup.Valid).Error; err != nil {
		return err
	}
	return nil
}
func (repo *CouponRepository) MakeCouponInvalid(id int) error {
	if err := repo.DB.Exec(`UPDATE coupons SET valid=false WHERE id=$1`, id).Error; err != nil {
		return err
	}
	return nil
}
func (repo *CouponRepository) FindCouponDiscount(couponID int) int {

	var coupon int

	err := repo.DB.Raw(`SELECT discount_rate FROM coupons WHERE id=$1`, couponID).Scan(&coupon).Error
	if err != nil {
		return 0
	}

	return coupon
}
