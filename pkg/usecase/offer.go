package usecase

import (
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
)

type OfferUseCase struct {
	repository interfaces.OfferRepository
}

func NewOfferusecase(repo interfaces.OfferRepository) usecaseInterfaces.OfferUseCase {
	return &OfferUseCase{
		repository: repo,
	}
}
func (off *OfferUseCase) AddNewOffer(model models.OfferMaking) error {
	if err := off.repository.AddNewOffer(model); err != nil {
		return err
	}
	return nil
}
func (off *OfferUseCase) MakeofferExpire(id int) error {
	if err := off.repository.MakeofferExpire(id); err != nil {
		return err
	}
	return nil
}
