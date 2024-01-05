package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	response "gokul.go/pkg/utils/response"
)

type CartHandler struct {
	usecase usecaseInterfaces.CartUseCase
}

func NewCartHandler(usecase usecaseInterfaces.CartUseCase) *CartHandler {
	return &CartHandler{
		usecase: usecase,
	}
}

// @Summary		Add To Cart
// @Description	Add products to carts  for the purchase
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			cart	body	models.AddToCart	true	"Add To Cart"
// @Security		BearerTokenAuth
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/home/add-to-cart [post]
func (i *CartHandler) AddToCart(c *gin.Context) {
	inventoryID := c.Query("inventory_id")
	inv_id, err := strconv.Atoi(inventoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "inv_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	userIDAny, ok := c.Get("id")
	if !ok {
		err := errors.New("cant get user id from context")
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the cart", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	userID, ok := userIDAny.(int)
	if !ok {
		err := errors.New("cant get user id from context")
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the cart", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	// if err := c.BindJSON(&model); err != nil {
	// 	errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errRes)
	// 	return
	// }
	if err := i.usecase.AddToCart(inv_id, userID); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added to cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Checkout section
// @Description	Add products to carts  for the purchase
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/check-out [get]
func (i *CartHandler) CheckOut(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := i.usecase.CheckOut(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}

//------mind checkout while doing ---------------commanding//
