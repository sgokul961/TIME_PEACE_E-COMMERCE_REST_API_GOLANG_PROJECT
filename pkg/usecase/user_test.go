package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"gokul.go/pkg/config"
	"gokul.go/pkg/mock"
	"gokul.go/pkg/utils/models"
)

func Test_User_UserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock implementations for the repositories
	userRepo := mock.NewMockUserRepository(ctrl)
	orderRepo := mock.NewMockOrderRepository(ctrl)
	otpRepo := mock.NewMockOtpRepository(ctrl)
	helper := mock.NewMockHelper(ctrl)
	cfg := config.Config{}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, helper, orderRepo)

	testData := map[string]struct {
		input          models.UserDetails
		stubDetails    func(mock.MockUserRepository, mock.MockOrderRepository, models.UserDetails, mock.MockHelper)
		expectedOutput models.TokenUsers
		expectedError  error
	}{
		"success": {
			input: models.UserDetails{
				Name:            "gokul",
				Email:           "gokul@gmail.com",
				Phone:           "6282288363",
				Password:        gomock.Any().String(),
				ConfirmPassWord: gomock.Any().String(),
			},
			stubDetails: func(userRepo mock.MockUserRepository, orderRepo mock.MockOrderRepository, signupData models.UserDetails, helper mock.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("1234").Times(1).Return(1, nil),
					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().UserSignUp(signupData, gomock.Any().String()).Times(1).Return(
						models.UserDeatilsResponse{
							Id:    1,
							Name:  signupData.Name,
							Email: signupData.Email,
							Phone: signupData.Phone,
						}, nil,
					),
					helper.EXPECT().GenerateTokenClients(models.UserDeatilsResponse{
						Id:    1,
						Name:  "gokul",
						Email: "gokul@gmail.com",
						Phone: "6282288363",
					}).Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(nil),
					orderRepo.EXPECT().CreateNewWallet(1).Times(1).Return(1, nil),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDeatilsResponse{
					Id:    1,
					Name:  "gokul",
					Email: "gokul@gmail.com",
					Phone: "6282288363",
				},
				Token: "abcdefghijklmop",
			},

			expectedError: nil,
		},
		"user already exist": {
			input: models.UserDetails{
				Name:            "gokul",
				Email:           "gokul@gmail.com",
				Phone:           "6282288363",
				Password:        gomock.Any().String(),
				ConfirmPassWord: gomock.Any().String(),
			},
			stubDetails: func(userRepo mock.MockUserRepository, orderRepo mock.MockOrderRepository, signupData models.UserDetails, helper mock.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(true),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("user already exist,sign in"),
		},
		"password missmatch": {
			input: models.UserDetails{
				Name:            "gokul",
				Email:           "gokul@gmail.com",
				Phone:           "6282288363",
				Password:        "goku",
				ConfirmPassWord: "guku",
			},
			stubDetails: func(userRepo mock.MockUserRepository, orderRepo mock.MockOrderRepository, signupData models.UserDetails, helper mock.MockHelper) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
				)
			},
			expectedOutput: models.TokenUsers{},
			expectedError:  errors.New("password dosnt match"),
		},
	}
	for name, test := range testData {
		fmt.Println("name", name)
		test.stubDetails(*userRepo, *orderRepo, test.input, *helper)
		tokenUsers, err := userUseCase.UserSignUp(test.input, "1234")

		assert.Equal(t, test.expectedOutput.Users.Id, tokenUsers.Users.Id)
		assert.Equal(t, test.expectedOutput.Users.Name, tokenUsers.Users.Name)
		assert.Equal(t, test.expectedOutput.Users.Email, tokenUsers.Users.Email)
		assert.Equal(t, test.expectedError, err)

	}

}
