package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	"gokul.go/pkg/utils/response"
)

type OfferHandler struct {
	usecase usecaseInterfaces.OfferUseCase
}

func NewOfferHandler(usecase usecaseInterfaces.OfferUseCase) *OfferHandler {
	return &OfferHandler{usecase: usecase}
}

// @Summary		Add Offer
// @Description	Admin can add new offers forspecified categories
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			offer	body	models.OfferMaking	true	"offer"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/offers/add [post]
func (off *OfferHandler) AddNewOffer(c *gin.Context) {
	var model models.OfferMaking

	if err := c.BindJSON(&model); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := off.usecase.AddNewOffer(model)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added offer", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (off *OfferHandler) MakeofferExpire(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := off.usecase.MakeofferExpire(id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "coupon cannot be invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	succRes := response.ClientResponse(http.StatusOK, "successffully made offer expire", nil, nil)
	c.JSON(http.StatusOK, succRes)
}
