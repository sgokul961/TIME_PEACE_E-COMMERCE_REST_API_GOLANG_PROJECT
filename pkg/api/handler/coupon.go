package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	"gokul.go/pkg/utils/response"
)

type CouponHAndler struct {
	usecase usecaseInterfaces.CouponUseCase
}

func NewCouponHandler(use usecaseInterfaces.CouponUseCase) *CouponHAndler {
	return &CouponHAndler{usecase: use}
}

// @Summary		Add Coupon
// @Description	Admin can add new coupons
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			coupon	body	models.Coupons	true	"coupon"
// @Security		BearerTokenAuth
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupon/create [post]
func (coup *CouponHAndler) CreateNewCoupon(c *gin.Context) {
	var coupon models.Coupons

	if err := c.BindJSON(&coupon); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feild provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := coup.usecase.Addcoupon(coupon)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added coupon", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Make Coupon ad invalid
// @Description	Admin can make the coupons as invalid so that users cannot use that particular coupon
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		BearerTokenAuth
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/coupons/delete [delete]
func (coup *CouponHAndler) MakeCouponInvalid(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := coup.usecase.MakeCouponInvalid(id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "coupon cannot be made invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully made coupon as invalid", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
