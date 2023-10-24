package interfaces

import "gokul.go/pkg/utils/models"

type OfferRepository interface {
	AddNewOffer(model models.OfferMaking) error
	MakeofferExpire(id int) error
}
