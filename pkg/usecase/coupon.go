package usecase

import (
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
)

type CouponUseCase struct {
	repository interfaces.CouponRepository
}

func NewCouponUseCase(repo interfaces.CouponRepository) usecaseInterfaces.CouponUseCase {
	return &CouponUseCase{
		repository: repo,
	}
}
func (coup *CouponUseCase) Addcoupon(coupon models.Coupons) error {
	if err := coup.repository.Addcoupon(coupon); err != nil {
		return err
	}
	return nil
}
func (coup *CouponUseCase) MakeCouponInvalid(id int) error {
	if err := coup.repository.MakeCouponInvalid(id); err != nil {
		return err
	}
	return nil
}
