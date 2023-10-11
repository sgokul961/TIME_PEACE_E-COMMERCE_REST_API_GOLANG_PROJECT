package usecaseInterfaces

import (
	"gokul.go/pkg/domain"
	"gokul.go/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory domain.Inventories) (models.InventoryResponse, error)
	UpdateInventory(productID uint, Stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(sku string) (domain.Inventories, error)
	ListProducts(page, count int) ([]domain.Inventories, error)
}
