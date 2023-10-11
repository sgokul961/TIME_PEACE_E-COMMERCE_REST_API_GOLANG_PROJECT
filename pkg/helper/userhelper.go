package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gokul.go/pkg/utils/models"
)

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {

// 	claims := &authCustomClaims{
// 		Id:    userID,
// 		Email: userEmail,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("132457689"))

// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func GenerateAccessToken(user models.UserDeatilsResponse) (string, error) {

// 	expirationTime := time.Now().Add(15 * time.Minute)
// 	tokenString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func GenerateRefreshToke(user models.UserDeatilsResponse) (string, error) {

// 	expirationTime := time.Now().Add(24 * 90 * time.Hour)
// 	tokeString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokeString, nil

// }

func GenerateTokenClients(user models.UserDeatilsResponse) (string, error) {
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
