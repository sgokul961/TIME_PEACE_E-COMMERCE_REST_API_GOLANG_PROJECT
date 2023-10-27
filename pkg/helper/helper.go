package helper

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	con "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/golang-jwt/jwt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
	"gokul.go/pkg/config"
	interfaces "gokul.go/pkg/helper/interface"
	"gokul.go/pkg/utils/models"
)

var client *twilio.RestClient

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type helper struct {
	cfg config.Config
}

// TwilioSetup implements interfaces.Helper.

func NewHelper(config config.Config) interfaces.Helper {
	return &helper{cfg: config}
}
func (h *helper) AddImageToS3(file *multipart.FileHeader) (string, error) {

	cfg, err := con.LoadDefaultConfig(context.TODO(), con.WithRegion("ap-south-1"))
	if err != nil {
		fmt.Println("configuration error:", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	f, openErr := file.Open()
	if openErr != nil {
		fmt.Println("opening error:", openErr)
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("time-peace"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		fmt.Println("uploading error:", uploadErr)
		return "", uploadErr
	}

	return result.Location, nil
}

// here we might need a change

func (h *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {

	claims := &authCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("comebuyjersey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (h *helper) TwilioSetup(username string, password string) {
	fmt.Println(username, password)
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

func (h *helper) TwilioSendOTP(phone string, serviceID string) (string, error) {
	fmt.Println("anybody here?")
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	fmt.Println(*params.To, *params.Channel)
	fmt.Println(serviceID)
	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {

		fmt.Println("CHECK CHECK")
		fmt.Println(err)
		return " ", err
	}

	return *resp.Sid, nil

}

func (h *helper) TwilioVerifyOTP(serviceID string, code string, phone string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)

	if err != nil {
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")

}
func (h *helper) GenerateTokenClients(user models.UserDeatilsResponse) (string, error) {
	claims := &authCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("comebuyjersey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (h *helper) GenerateRefferalCode() (string, error) {
	// Calculate the required number of random bytes
	byteLength := (5 * 5) / 8

	// Generate a random byte array
	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base32
	encoder := base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding)
	encoded := encoder.EncodeToString(randomBytes)

	// Trim any additional characters to match the desired length
	encoded = encoded[:5]

	return encoded, nil
}
