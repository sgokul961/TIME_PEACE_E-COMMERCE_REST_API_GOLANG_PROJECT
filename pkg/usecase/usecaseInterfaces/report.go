package usecaseInterfaces

import (
	"time"

	"gokul.go/pkg/domain"
)

type SalesUseCase interface {
	GetMonthlySalesReport(year int, month time.Month) ([]domain.Order, error)
}
