package usecase

import (
	"fmt"
	"strconv"

	"github.com/razorpay/razorpay-go"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
)

type paymentusecase struct {
	repository interfaces.PayementRepository
}

func NewPaymentUseCase(repo interfaces.PayementRepository) usecaseInterfaces.PaymentUseCase {
	return &paymentusecase{
		repository: repo,
	}
}
func (p *paymentusecase) MakePaymentRazorPay(orderID string, userID string) (models.OrderPaymentDetails, error) {

	var orderDetails models.OrderPaymentDetails

	//get order id

	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = newid

	fmt.Println("1")

	//get userid
	newuserid, err := strconv.Atoi(userID)

	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.UserID = newuserid

	fmt.Println("2")

	//get user payment id

	fmt.Println("before find username")
	username, err := p.repository.FindUserName(newuserid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.Username = username

	fmt.Println("after find username")

	fmt.Println("3")

	//get totall

	newfinal, err := p.repository.FindPrice(newid)

	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.FinalPrice = newfinal

	fmt.Println("4")

	client := razorpay.NewClient("rzp_test_pfmFeCViv6CU5K", "TWCh1tyyZZsIxjYSOmmRrLLg")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	fmt.Println("5")
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.OrderPaymentDetails{}, nil
	}
	razoPayOrderID := body["id"].(string)

	orderDetails.Razor_id = razoPayOrderID

	fmt.Println("5")

	return orderDetails, nil

}
func (p *paymentusecase) VerifyPayment(paymentID string, razorID string, orderID string) error {

	err := p.repository.UpdatePaymentDetails(orderID, paymentID, razorID)
	if err != nil {
		return err
	}

	return nil
}
func (p *paymentusecase) UseWallwt(OrderID string, userID string) (models.OrderPaymentDetails, error) {

	var orderDetails models.OrderPaymentDetails

	//get orderID

	newid, err := strconv.Atoi(OrderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = newid

	//get userID

	newuserID, err := strconv.Atoi(userID)

	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.UserID = newuserID

	//get username

	username, err := p.repository.FindUserName(newuserID)

	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.Username = username

	//get total
	newfinal, er := p.repository.FindPrice(newid)

	if err != nil {
		return models.OrderPaymentDetails{}, er
	}

	orderDetails.FinalPrice = newfinal

	client := razorpay.NewClient("rzp_test_YdhpOaBfGPNYLu", "lP5LQ9No01YVmg7GG8WRJeZ5")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.OrderPaymentDetails{}, nil
	}
	razorPayOrderID := body["id"].(string)

	orderDetails.Razor_id = razorPayOrderID

	return orderDetails, nil

}
