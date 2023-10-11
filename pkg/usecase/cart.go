package usecase

import (
	"errors"

	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
	userusecase         usecaseInterfaces.UserUseCase
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository, userusecase usecaseInterfaces.UserUseCase) usecaseInterfaces.CartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
		userusecase:         userusecase, ///check here for errr
	}
}

func (i *cartUseCase) AddToCart(user_id, inventory_id int) error {
	//check the quqntity of product

	stock, err := i.inventoryRepository.CheckStock(inventory_id)
	if err != nil {
		return err
	}
	//if available call userRepository
	if stock <= 0 {
		return errors.New("out of stock")
	}

	//find user cart id

	cart_id, err := i.repo.GetCartId(user_id)
	if err != nil {
		return errors.New("error getting user cart")

	}
	//if user has no existing cart create a new cart

	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(user_id)
		if err != nil {
			return errors.New("cannot create cart for user")
		}
	}
	//add product to line items

	if err := i.repo.AddLineItems(cart_id, inventory_id); err != nil {
		return errors.New("error in adding products")
	}
	return nil
}

func (i *cartUseCase) CheckOut(id int) (models.CheckOut, error) {

	address, err := i.repo.GetAddresses(id)

	if err != nil {
		return models.CheckOut{}, err
	}
	payment, err := i.repo.GetPaymentOptions()
	if err != nil {
		return models.CheckOut{}, err
	}

	products, err := i.userusecase.GetCart(id)

	if err != nil {
		return models.CheckOut{}, err
	}
	var price float64
	for _, v := range products {
		price = price + v.Totel
	}
	var checkout models.CheckOut

	checkout.Addresses = address
	checkout.Products = products
	checkout.PaymentMethord = payment
	checkout.TotelPrice = price

	return checkout, err
}
