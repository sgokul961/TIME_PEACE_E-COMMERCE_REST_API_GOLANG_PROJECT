package usecase

import (
	"time"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
)

type SalesUseCase struct {
	salesrepo interfaces.SalesRepository
}

func NewSalesUseCase(repo interfaces.SalesRepository) usecaseInterfaces.SalesUseCase {
	return &SalesUseCase{
		salesrepo: repo,
	}
}
func (u *SalesUseCase) GetMonthlySalesReport(year int, month time.Month) ([]domain.Order, error) {
	orders, err := u.salesrepo.GetMonthlySalesReport(year, month)

	if err != nil {
		return nil, err
	}
	return orders, nil
}
