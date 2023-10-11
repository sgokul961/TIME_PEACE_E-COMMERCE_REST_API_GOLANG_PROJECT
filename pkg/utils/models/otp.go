package models

type OTPData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
}
type VarifyData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}
type PaymentMethod struct {
	ID          uint   `json:"id"`
	PaymentName string `json:"payment_name"`
}
