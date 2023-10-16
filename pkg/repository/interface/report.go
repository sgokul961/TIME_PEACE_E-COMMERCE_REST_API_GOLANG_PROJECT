package interfaces

import (
	"time"

	"gokul.go/pkg/domain"
)

type SalesRepository interface {
	GetMonthlySalesReport(year int, month time.Month) ([]domain.Order, error)
}
