package usecaseInterfaces

import (
	"gokul.go/pkg/domain"
	"gokul.go/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	BlockUser(id string) error
	UnblockUser(id string) error
	GetUsers(page int, count int) ([]models.UserDetailsAdmin, error)
}
