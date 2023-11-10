package interfaces

import (
	"gokul.go/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory models.AddInventories, images []string) (models.InventoryResponse, error)
	CheckInventory(pid uint) (bool, error)
	UpdateInventory(pid uint, stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ListProducts(page int, count int) ([]models.InvResponse, error)
	CheckStock(inventory_id int) (int, error)

	GetInventory(inv_id int) (models.InvResponse, error)
	GetImages(inv_id int) ([]string, error)
}
