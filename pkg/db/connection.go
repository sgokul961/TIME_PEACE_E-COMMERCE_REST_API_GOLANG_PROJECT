package db

import (
	"fmt"

	"gokul.go/pkg/config"
	"gokul.go/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	// if dbErr != nil {
	// 	return nil, dbErr
	// }

	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(domain.Inventories{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(domain.LineItems{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.Offer{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderItem{})
	db.AutoMigrate(domain.Coupons{})
	db.AutoMigrate(domain.Offer{})
	db.AutoMigrate(domain.Wallet{})

	CheckAndCreateAdmin(db)

	return db, dbErr
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "patek"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		admin := domain.Admin{
			ID:       3,
			Name:     "timepeace",
			Email:    "gokul@gmail.com",
			Password: string(hashedPassword),
		}
		db.Create(&admin)
	}
}
