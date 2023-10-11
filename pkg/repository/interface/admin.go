package interfaces

import (
	"gokul.go/pkg/domain"
	"gokul.go/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserById(id string) (domain.Users, error)
	UpdateBlockUserID(user domain.Users) error
	GetUsers(page int, count int) ([]models.UserDetailsAdmin, error)
}
