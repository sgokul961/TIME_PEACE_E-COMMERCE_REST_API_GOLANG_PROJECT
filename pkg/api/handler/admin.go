package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	response "gokul.go/pkg/utils/response"
)

type AdminHandler struct {
	adminUseCase usecaseInterfaces.AdminUseCase
}

func NewAdminHandler(usecase usecaseInterfaces.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase}
}

// @Summary		Admin Login
// @Description	Login handler for timepeace admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AdminLogin	true	"Admin login details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/adminlogin [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) {

	var adminDetails models.AdminLogin

	fmt.Println("it is here ")
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in proper format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Block User
// @Description	using this handler admins can block an user
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/users/block [post]
func (ad *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Query("id")
	err := ad.adminUseCase.BlockUser(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully blocked user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		UnBlock an existing user
// @Description	UnBlock user
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security BearerTokenAuth
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/users/unblock [POST]
func (ad *AdminHandler) UnblockUser(c *gin.Context) {

	id := c.Query("id")
	err := ad.adminUseCase.UnblockUser(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user cant be unblocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Get Users
// @Description	Retrieve users with pagination
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security BearerTokenAuth
// @Param			page	query		int	true	"Page number"
// @param           count    query      int     true    "count"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/users/getusers [get]
func (ad *AdminHandler) GetUsers(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page num is not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	count, err := strconv.Atoi(c.Query("count"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	users, err := ad.adminUseCase.GetUsers(page, count)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully retrive the users", users, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Orderstatus
// @Description	Orderstatus handler for timepeace admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security BearerTokenAuth
// @Param			order_status	query  string	true	"order_status"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/orders/status [get]
func (i *AdminHandler) Orderstatus(ctx *gin.Context) {

	status := ctx.Query("order_status")

	pendingorder, err := i.adminUseCase.Orderstatus(status)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, pendingorder)

}
