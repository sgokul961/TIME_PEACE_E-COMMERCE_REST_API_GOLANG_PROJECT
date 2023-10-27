package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	"gokul.go/pkg/utils/response"
)

type OtpHandler struct {
	otpUseCase usecaseInterfaces.OtpUseCase
}

func NewOtpHandler(useCase usecaseInterfaces.OtpUseCase) *OtpHandler {
	return &OtpHandler{otpUseCase: useCase}
}

func (ot *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData

	if err := c.BindJSON(&phone); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}
	err := ot.otpUseCase.SendOTP(phone.PhoneNumber)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "OTP sent Successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
func (ot *OtpHandler) VerifyOTP(c *gin.Context) {
	fmt.Println(1)
	var code models.VarifyData
	if err := c.BindJSON(&code); err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	fmt.Println(2)

	users, err := ot.otpUseCase.VerifyOTP(code)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not varify otp", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println(3)
	successRes := response.ClientResponse(http.StatusOK, "Successfully verifyed OTP", users, nil)
	c.JSON(http.StatusOK, successRes)
}
