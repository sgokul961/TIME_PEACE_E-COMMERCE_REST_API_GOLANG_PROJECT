package interfaces

import "gokul.go/pkg/utils/models"

type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDeatilsUsingPhone(phone string) (models.UserDeatilsResponse, error)
}
