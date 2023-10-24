package usecaseInterfaces

import "gokul.go/pkg/utils/models"

type CouponUseCase interface {
	Addcoupon(coupon models.Coupons) error
	MakeCouponInvalid(id int) error
}
