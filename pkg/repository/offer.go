package repository

import (
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/utils/models"
	"gorm.io/gorm"
)

type OfferRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(db *gorm.DB) interfaces.OfferRepository {
	return &OfferRepository{
		DB: db,
	}
}
func (repo *OfferRepository) AddNewOffer(model models.OfferMaking) error {
	if err := repo.DB.Exec(`INSERT INTO offers(category_id,discount_rate)VALUES ($1,$2)`, model.CategoryID, model.Discount).Error; err != nil {
		return err
	}
	return nil
}
func (repo *OfferRepository) MakeofferExpire(id int) error {
	if err := repo.DB.Exec(`UPDATE offers SET valid=false WHERE id=$1`, id).Error; err != nil {
		return err

	}
	return nil

}
