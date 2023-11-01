package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	"gokul.go/pkg/utils/response"
)

type UserHandler struct {
	userUsecase usecaseInterfaces.UserUseCase
}

type Response struct {
	Name    string `copier:"must"`
	Surname string `copier:"must"`
	ID      uint   `copier:"must"`
}

func NewUserHandler(usecase usecaseInterfaces.UserUseCase) *UserHandler {
	return &UserHandler{userUsecase: usecase}
}

// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			signup  body  models.UserDetails  true	"signup"
// @Success		200	{object}	response.Response{} ""
// @Failure		500	{object}	response.Response{} ""
// @Router			/user/signup [post]
func (u *UserHandler) UserSignUp(c *gin.Context) {

	var user models.UserDetails

	//bind user details to the struct

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}

	//checking wether the data sent by the user has all the correct constrains specified by the User struct

	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constrains not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	//if the user wants to mention the referral code of other user
	ref := c.Query("reference")

	//bussiness logic goes inside this function
	userCreated, err := u.userUsecase.UserSignUp(user, ref) //check here

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User coudnt sign up", nil, err.Error())
		c.JSON(http.StatusCreated, errRes)

		return

	}
	successRes := response.ClientResponse(http.StatusCreated, "user successfully signed up", userCreated, nil)

	c.JSON(http.StatusCreated, successRes)

}

// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			login  body  models.UserLogin  true	"login"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/login [post]
func (u *UserHandler) LoginHandler(c *gin.Context) {

	var user models.UserLoign

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feild provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	err := validator.New().Struct(user)
	if err != nil {  
		errRes := response.ClientResponse(http.StatusBadRequest, "constrains not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user_details, err := u.userUsecase.LoginHandler(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user could not be logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "user successfully logged in", user_details, nil)
	c.JSON(http.StatusOK, successRes)

}
func (i *UserHandler) AddAddress(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check path parameater", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var address models.AddAddress

	if err := c.BindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	if err := i.userUsecase.AddAddress(id, address); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully added address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *UserHandler) GetAddress(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check your id agin", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	addresses, err := i.userUsecase.GetAddress(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", addresses, nil)
	c.JSON(http.StatusOK, successRes)

}
func (i *UserHandler) GetUserDetails(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	details, err := i.userUsecase.GetUserDetails(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	successRes := response.ClientResponse(http.StatusOK, "Successful got all records", details, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *UserHandler) ChangePassword(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check path parameater", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var ChangePassword models.ChangePassword

	if err := c.BindJSON(&ChangePassword); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := i.userUsecase.ChangePassword(id, ChangePassword.Oldpassword, ChangePassword.Password, ChangePassword.Repassword); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	successRes := response.ClientResponse(http.StatusOK, "password changed successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *UserHandler) ForgotPasswordSend(c *gin.Context) {

	var model models.ForgotPasswordSend

	if err := c.BindJSON(&model); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := i.userUsecase.ForgotPasswordSend(model.Phone)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not send OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *UserHandler) GetCart(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check parameaters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	products, err := i.userUsecase.GetCart(id)

	if err != nil {
		errrRes := response.ClientResponse(http.StatusBadRequest, "could not retrive cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errrRes)
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all products in cart", products, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *UserHandler) RemoveFromCart(c *gin.Context) {

	cart_id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errREs := response.ClientResponse(http.StatusBadRequest, "check parameaters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errREs)
		return
	}
	inventory_id, err := strconv.Atoi(c.Query("inventory"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cheeck parameaters proprrly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := i.userUsecase.RemoveFromCart(cart_id, inventory_id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully removed the product", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (i *UserHandler) UpdateQuantityAdd(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check parameaters proprly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	inventory_id, err := strconv.Atoi(c.Query("inventory_id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check the parameaters", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := i.userUsecase.UpdateQuantityAdd(id, inventory_id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusOK, "succesfully added quantity", nil, nil)
	c.JSON(http.StatusOK, succesRes)

}
func (i *UserHandler) UpdateQuantityMinus(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check parameaters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	inventory_id, err := strconv.Atoi(c.Query("inventory_id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check parameaters", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := i.userUsecase.UpdateQuantityMinus(id, inventory_id); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "coudnt add the quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully removed item", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (i *UserHandler) VarifyForgotPasswordAndChange(c *gin.Context) {
	var models models.ForgotVarify

	if err := c.BindJSON(&models); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	err := i.userUsecase.VarifyForgotPasswordAndChange(models)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not varify otp", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully cahnged the password", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *UserHandler) GetMyReferanceLink(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "check paramesters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	link, err := i.userUsecase.GetMyReferanceLink(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrive refferal link", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all product in cart", link, nil)
	c.JSON(http.StatusOK, successRes)
}
