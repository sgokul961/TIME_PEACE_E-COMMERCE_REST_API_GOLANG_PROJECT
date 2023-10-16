package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	response "gokul.go/pkg/utils/response"
)

type OrdeHandler struct {
	orderUseCase usecaseInterfaces.OrderUseCase
}

func NewOrderHandler(useCase usecaseInterfaces.OrderUseCase) *OrdeHandler {
	return &OrdeHandler{orderUseCase: useCase}
}
func (i *OrdeHandler) GetOrders(c *gin.Context) {

	idString := c.Query("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orders, err := i.orderUseCase.GetOrders(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *OrdeHandler) OrderItemsFromCart(c *gin.Context) {

	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := i.orderUseCase.OrderItemsFromCart(order.UserID, order.AddressID, order.PaymentMethodID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully made the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *OrdeHandler) CancelOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "convertion to integer not posssible", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := i.orderUseCase.CancelOrder(id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided in wrong format ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully canceled the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *OrdeHandler) AdminOrders(c *gin.Context) {

	orders, err := i.orderUseCase.AdminOrders()

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all record", orders, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *OrdeHandler) EditOrderStatus(c *gin.Context) {
	status := c.Query("status")
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion into integr not possible", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := i.orderUseCase.EditOrderStatus(status, id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully edited the order status", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
