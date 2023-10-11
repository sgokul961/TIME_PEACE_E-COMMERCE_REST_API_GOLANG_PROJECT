package interfaces

import (
	"gokul.go/pkg/domain"
	"gokul.go/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory domain.Inventories) (models.InventoryResponse, error)
	CheckInventory(pid uint) (bool, error)
	UpdateInventory(pid uint, stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(id string) (domain.Inventories, error)
	ListProducts(page int, count int) ([]domain.Inventories, error)
	CheckStock(inventory_id int) (int, error)
}
