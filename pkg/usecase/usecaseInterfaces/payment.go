package usecaseInterfaces

import "gokul.go/pkg/utils/models"

type PaymentUseCase interface {
	MakePaymentRazorPay(orderID string, userID string) (models.OrderPaymentDetails, error)
	VerifyPayment(paymentID string, razorID string, orderID string) error
	UseWallwt(orderID string, userID string) (models.OrderPaymentDetails, error)
}
