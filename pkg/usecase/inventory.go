package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"

	"gokul.go/pkg/domain"
	helper_interface "gokul.go/pkg/helper/interface"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
)

type inventoryUseCase struct {
	repository interfaces.InventoryRepository
	helper     helper_interface.Helper
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, helper helper_interface.Helper) usecaseInterfaces.InventoryUseCase {

	return &inventoryUseCase{
		repository: repo,
		helper:     helper}
}
func (i *inventoryUseCase) AddInventory(inventory models.AddInventories, image *multipart.FileHeader) (models.InventoryResponse, error) {
	// InventoryResponse, err := i.repository.AddInventory(inventory)
	// if err != nil {
	// 	return models.InventoryResponse{}, err
	// }
	// return InventoryResponse, nil
	url, err := i.helper.AddImageToS3(image)

	if err != nil {
		return models.InventoryResponse{}, err
	}
	//send url in databe
	InventoryResponse, err := i.repository.AddInventory(inventory, url)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return InventoryResponse, nil

}

func (i *inventoryUseCase) UpdateInventory(pid uint, stock int) (models.InventoryResponse, error) {

	result, err := i.repository.CheckInventory(pid)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	if !result {
		fmt.Println("2")
		return models.InventoryResponse{}, errors.New("there is no inventory as you mentioned")
	}
	newcat, err := i.repository.UpdateInventory(pid, stock)
	if err != nil {
		fmt.Println("3")
		return models.InventoryResponse{}, err
	}
	return newcat, err
}
func (i *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := i.repository.DeleteInventory(inventoryID)

	if err != nil {
		return err
	}
	return nil
}
func (i *inventoryUseCase) ShowIndividualProducts(id string) (domain.Inventories, error) {

	product, err := i.repository.ShowIndividualProducts(id)

	if err != nil {
		return domain.Inventories{}, err
	}
	return product, nil
}
func (i *inventoryUseCase) ListProducts(page int, count int) ([]domain.Inventories, error) {
	productDetails, err := i.repository.ListProducts(page, count)
	if err != nil {
		return []domain.Inventories{}, err
	}
	return productDetails, nil
}
