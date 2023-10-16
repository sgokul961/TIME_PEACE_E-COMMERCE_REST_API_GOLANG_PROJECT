// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"gokul.go/pkg/api"
	"gokul.go/pkg/api/handler"
	"gokul.go/pkg/config"
	"gokul.go/pkg/db"
	"gokul.go/pkg/repository"
	"gokul.go/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	userRepository := repository.NewUserRepository(gormDB)
	otpRepository := repository.NewOtpRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cfg, otpRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository)
	otpHandler := handler.NewOtpHandler(otpUseCase)
	inventoryRepository := repository.NewInventoryRepository(gormDB)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)
	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, inventoryRepository, userUseCase)
	cartHandler := handler.NewCartHandler(cartUseCase)
	orderRepository := repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, userUseCase)
	ordeHandler := handler.NewOrderHandler(orderUseCase)
	payementRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(payementRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)
	salesRepository := repository.NewSalesRepository(gormDB)
	salesUseCase := usecase.NewSalesUseCase(salesRepository)
	salesHandler := handler.NewSalesHandler(salesUseCase)
	serverHTTP := http.NewServerHTTP(adminHandler, userHandler, otpHandler, inventoryHandler, categoryHandler, cartHandler, ordeHandler, paymentHandler, salesHandler)
	return serverHTTP, nil
}
