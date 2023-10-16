//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "gokul.go/pkg/api"
	"gokul.go/pkg/api/handler"
	"gokul.go/pkg/config"
	"gokul.go/pkg/db"
	"gokul.go/pkg/repository"
	"gokul.go/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
		handler.NewUserHandler,

		repository.NewCategoryRepository,
		usecase.NewCategoryUseCase,
		handler.NewCategoryHandler,

		repository.NewInventoryRepository,
		usecase.NewInventoryUseCase,
		handler.NewInventoryHandler,

		repository.NewOtpRepository,
		usecase.NewOtpUseCase,
		handler.NewOtpHandler,

		repository.NewCartRepository,
		usecase.NewCartUseCase,
		handler.NewCartHandler,

		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		handler.NewOrderHandler,

		repository.NewPaymentRepository,
		usecase.NewPaymentUseCase,
		handler.NewPaymentHandler,

		repository.NewSalesRepository,
		usecase.NewSalesUseCase,
		handler.NewSalesHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}

//mistyped servmux and manually changed to serverHTTP
