package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/response"
)

type PaymentHandler struct {
	usecase usecaseInterfaces.PaymentUseCase
}

func NewPaymentHandler(use usecaseInterfaces.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{usecase: use}
}
func (p *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {
	OrderID := c.Query("id")

	userID := c.Query("user_id")

	orderDetail, err := p.usecase.MakePaymentRazorPay(OrderID, userID)

	//fmt.Println("orderdetails:", orderDetail)
	//fmt.Println("8")

	if err != nil {

		//fmt.Println("9")
		
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not genarate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	fmt.Println("10")

	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}

func (p *PaymentHandler) VerifyPayment(c *gin.Context) {
	OrderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")

	err := p.usecase.VerifyPayment(paymentID, razorID, OrderID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not update error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully ", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (p *PaymentHandler) MakePaymentFromWallet(c *gin.Context) {
	orderID := c.Query("order_id")
	userID := c.Query("user_id")

	orderDetail, err := p.usecase.UseWallwt(orderID, userID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "could not make payment from wallet", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}
