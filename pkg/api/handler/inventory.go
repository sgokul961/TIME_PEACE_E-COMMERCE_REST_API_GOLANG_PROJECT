package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

// @Summary		Add Inventory
// @Description	Admin can add new  products
// @Tags			Admin
// @Accept			multipart/form-data
// @Produce		    json
// @Param			category_id		formData	string	true	"category_id"
// @Param			product_name	formData	string	true	"product_name"
// @Param			size		formData	string	true	"size"
// @Param			price	formData	string	true	"price"
// @Param			stock		formData	string	true	"stock"
// @Param           image      formData     file   true   "images"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/add [post]
func (i *InventoryHandler) AddInventory(c *gin.Context) {

	var inventory models.AddInventories

	categoryID, err := strconv.Atoi(c.Request.FormValue("category_id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	product_name := c.Request.FormValue("product_name")
	size := c.Request.FormValue("size")

	p, err := strconv.ParseFloat(c.Request.FormValue("price"), 64)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	price := float64(p)

	stock, err := strconv.Atoi(c.Request.FormValue("stock"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	inventory.CategoryID = categoryID
	inventory.ProductName = product_name
	inventory.Size = size
	inventory.Price = price
	inventory.Stock = stock

	// file, err := c.MultipartForm()
	// if err != nil {
	// 	errRes := response.ClientResponse(http.StatusBadRequest, "retrieving image from form error", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errRes)
	// 	return
	// }
	// InventoryResponse, err := i.InventoryUseCase.AddInventory(inventory,file)
	// if err != nil {
	// 	errRes := response.ClientResponse(http.StatusBadRequest, "could not add the inventory", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errRes)
	// 	return
	// }
	// successRes := response.ClientResponse(http.StatusOK, "successfully added inventory", InventoryResponse, nil)
	// c.JSON(http.StatusOK, successRes)
	// fmt.Println("8")

	err = c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	files := c.Request.MultipartForm.File["images"]

	InventoryResponse, err := i.InventoryUseCase.AddInventory(inventory, files)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added inventory", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Update Stock
// @Description	Admin can update stock of the inventories
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			add-stock	body	models.InventoryUpdate	true	"update stock"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/update [put]
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

// @Summary		Delete Inventory
// @Description	Admin can delete a product
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/delete [delete]
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
func (i *InventoryHandler) GetIndividualProducts(c *gin.Context) {

	invID, err := strconv.Atoi(c.Query("inv_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "inventorys in the page not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	inventory, err := i.InventoryUseCase.GetIndividualProducts(invID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)

		return

	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", inventory, nil)
	c.JSON(http.StatusOK, successRes)

}
