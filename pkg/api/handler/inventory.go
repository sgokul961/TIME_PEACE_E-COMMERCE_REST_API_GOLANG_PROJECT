package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/domain"
	"gokul.go/pkg/usecase/usecaseInterfaces"

	"gokul.go/pkg/utils/models"
	"gokul.go/pkg/utils/response"
)

type InventoryHandler struct {
	InventoryUseCase usecaseInterfaces.InventoryUseCase
}

func NewInventoryHandler(usecase usecaseInterfaces.InventoryUseCase) *InventoryHandler {
	return &InventoryHandler{InventoryUseCase: usecase}
}

func (i *InventoryHandler) AddInventory(c *gin.Context) {

	var Inventory domain.Inventories

	if err := c.BindJSON(&Inventory); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	InventoryResponse, err := i.InventoryUseCase.AddInventory(Inventory)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added inventory", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *InventoryHandler) UpdateInventory(c *gin.Context) {

	var p models.InventoryUpdate

	if err := c.BindJSON(&p); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	a, err := i.InventoryUseCase.UpdateInventory(p.Productid, p.Stock)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not update the inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully updated the inventory stock", a, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *InventoryHandler) DeleteInventory(c *gin.Context) {

	inventoryID := c.Query("id")
	fmt.Println("inventoryID is ", inventoryID)

	err := i.InventoryUseCase.DeleteInventory(inventoryID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully Deleted the inventory", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *InventoryHandler) ShowIndividualProducts(c *gin.Context) {
	id := c.Query("id")
	product, err := i.InventoryUseCase.ShowIndividualProducts(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "path variables are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	successRes := response.ClientResponse(http.StatusOK, "product details retrived successfully", product, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *InventoryHandler) ListProducts(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	count, err := strconv.Atoi(c.Query("count"))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in the page not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	products, err := i.InventoryUseCase.ListProducts(page, count)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)

		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)

}
