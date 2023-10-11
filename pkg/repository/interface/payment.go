package interfaces

type PayementRepository interface {
	FindUserName(user_id int) (string, error)
	FindPrice(order_id int) (float64, error)
	UpdatePaymentDetails(orderID, paymentID, razorID string) error
}
