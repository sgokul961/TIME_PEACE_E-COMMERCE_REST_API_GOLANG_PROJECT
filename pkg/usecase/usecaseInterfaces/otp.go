package usecaseInterfaces

import "gokul.go/pkg/utils/models"

type OtpUseCase interface {
	VerifyOTP(code models.VarifyData) (models.TokenUsers, error)
	SendOTP(phone string) error
}
