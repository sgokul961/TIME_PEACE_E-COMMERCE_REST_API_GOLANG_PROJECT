package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"

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
		helper:     helper,
	}
}
func (i *inventoryUseCase) AddInventory(inventory models.AddInventories, image []*multipart.FileHeader) (models.InventoryResponse, error) {
	// InventoryResponse, err := i.repository.AddInventory(inventory)
	// if err != nil {
	// 	return models.InventoryResponse{}, err
	// }
	// return InventoryResponse, nil
	/*url, err := i.helper.AddImageToS3(image)*/

	/*if err != nil {
	return models.InventoryResponse{}, err*/
	//}
	//send url in databe

	//images := image

	var imageUrls []string

	for _, fileHeader := range image {

		Uploadurl, err := i.helper.AddImageToS3(fileHeader)
		if err != nil {
			return models.InventoryResponse{}, err
		}
		imageUrls = append(imageUrls, Uploadurl)
	}
	InventoryResponse, err := i.repository.AddInventory(inventory, imageUrls)
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

func (i *inventoryUseCase) ListProducts(page int, count int) ([]models.InvResponse, error) {
	productDetails, err := i.repository.ListProducts(page, count)
	if err != nil {
		return []models.InvResponse{}, err
	}
	updatedproductDetails := make([]models.InvResponse, 0)
	for _, k := range productDetails {
		img, err := i.repository.GetImages(int(k.ID))
		if err != nil {
			return nil, err
		}
		k.Image = img
		updatedproductDetails = append(updatedproductDetails, k)
	}

	return updatedproductDetails, nil

}
func (i *inventoryUseCase) GetIndividualProducts(inv_id int) (models.InvResponse, error) {

	inv, err := i.repository.GetInventory(inv_id)
	if err != nil {
		return models.InvResponse{}, err
	}
	imgURL, err := i.repository.GetImages(inv_id)
	if err != nil {
		return models.InvResponse{}, err
	}
	inv.Image = imgURL
	return inv, nil

}
