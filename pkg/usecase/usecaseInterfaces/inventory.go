package usecaseInterfaces

import (
	"mime/multipart"

	"gokul.go/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory models.AddInventories, images []*multipart.FileHeader) (models.InventoryResponse, error)
	UpdateInventory(productID uint, Stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ListProducts(page, count int) ([]models.InvResponse, error)
	GetIndividualProducts(inv_id int) (models.InvResponse, error)
}
