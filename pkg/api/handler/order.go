package handler

import (
	"errors"
	"fmt"
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

	idstring := c.Query("coupon-id")
	if idstring == "" {
		idstring = "0"
	}
	couponId, err := strconv.Atoi(idstring)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "coupon id problem", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println("check 4")
	userID, ok := c.Get("id")
	if !ok {
		err := errors.New("cant get user id from context")
		errRes := response.ClientResponse(http.StatusBadRequest, "could not order from cart", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id, ok := userID.(int)
	if !ok {
		err := errors.New("user id is not of type int")
		errRes := response.ClientResponse(http.StatusBadRequest, "could not order from cart", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println("check 3")
	if err := i.orderUseCase.OrderItemsFromCart(id, order.AddressID, order.PaymentMethodID, couponId); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println("check here")

	succRes := response.ClientResponse(http.StatusOK, "successfully placed the order", nil, nil)
	c.JSON(http.StatusOK, succRes)
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

func (h *OrdeHandler) GenerateInvoice(c *gin.Context) {
	//orderIDParam := c.Param("orderID")

	// orderID, err := strconv.ParseUint(orderIDParam, 10, 24)
	orderID, err := strconv.Atoi(c.Query("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid orderID"})
		return
	}
	pdf, err := h.orderUseCase.GenerateInvoice(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to genarateinvoice"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=invoice.pdf")

	// Generate a temporary file path for the PDF
	pdfFilePath := "salesReport/file.pdf"

	// Save the PDF to the temporary file path
	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		response := response.ClientResponse(500, "Failed to generate PDF", nil, err.Error())
		c.JSON(500, response)
		return
	}

	// Set the appropriate headers for the file download
	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")

	// Serve the PDF file for download
	c.File(pdfFilePath)

	// Set Content-Type header to application/pdf
	c.Header("Content-Type", "application/pdf")

	// Write PDF data to the response writer
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to serve invoice"})
		return
	}
	c.Status(http.StatusOK)

}
