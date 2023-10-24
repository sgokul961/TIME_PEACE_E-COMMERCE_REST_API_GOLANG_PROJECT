package usecaseInterfaces

import "gokul.go/pkg/utils/models"

type OfferUseCase interface {
	AddNewOffer(model models.OfferMaking) error
	MakeofferExpire(id int) error
}
