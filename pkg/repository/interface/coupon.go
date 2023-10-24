package interfaces

import "gokul.go/pkg/utils/models"

type CouponRepository interface {
	Addcoupon(models.Coupons) error
	MakeCouponInvalid(id int) error
	FindCouponDiscount(couponID int) int
}
