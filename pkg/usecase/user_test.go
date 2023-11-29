package usecase

// import (
// 	"testing"

// 	"github.com/go-playground/assert/v2"
// 	"github.com/golang/mock/gomock"
// 	"gokul.go/mocks"
// 	"gokul.go/pkg/config"
// 	"gokul.go/pkg/utils/models"
// )

// func Test_User_UserSignUp(t *testing.T) {
// 	cntrl := gomock.NewController(t)
// 	// Create mock implementations for the repositories

// 	userRepo := mocks.NewMockUserRepository(cntrl)
// 	orderRepo := mocks.NewMockOrderRepository(cntrl)
// 	otpRepo := mocks.NewMockOtpRepository(cntrl)
// 	helper := mocks.NewMockHelper(cntrl)
// 	cfg := config.Config{}

// 	type args struct {
// 		user models.UserDetails
// 		ref  string
// 	}

// 	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, orderRepo, helper)
// 	testData := map[string]struct {
// 		input          args
// 		StubDetails    func(mocks.MockUserRepository, mocks.MockHelper, args)
// 		expectedOutput models.TokenUsers
// 		expectedError  error
// 	}{
// 		"success": {
// 			input: args{user: models.UserDetails{
// 				Name:            "gokul",
// 				Email:           "gokul@gmail.com",
// 				Phone:           "6282288363",
// 				Password:        "hffj",
// 				ConfirmPassWord: "hffj",
// 			}, ref: "gok",
// 			},

// 			StubDetails: func(userRepo mocks.MockUserRepository, helper mocks.MockHelper, signupData args) {
// 				gomock.InOrder(
// 					userRepo.EXPECT().CheckUserAvailability(signupData.user.Email).Times(1).Return(false),
// 					userRepo.EXPECT().FindUserFromReference(signupData.ref).Times(1).Return(1, nil),
// 					helper.EXPECT().GenerateRefferalCode().Times(1).Return("goku", nil),
// 					userRepo.EXPECT().UserSignUp(signupData.user, "goku").Times(1).Return(
// 						models.UserDeatilsResponse{
// 							Id:    2,
// 							Name:  signupData.user.Name,
// 							Email: signupData.user.Email,
// 							Phone: signupData.user.Phone,
// 						}, nil),
// 					helper.EXPECT().GenerateTokenClients(models.UserDeatilsResponse{
// 						Id:    2,
// 						Name:  signupData.user.Name,
// 						Email: signupData.user.Email,
// 						Phone: signupData.user.Phone,
// 					}).Times(1).Return("ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs", nil),
// 					userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(nil),
// 					orderRepo.EXPECT().CreateNewWallet(2).Times(1).Return(1, nil),
// 				)
// 			},

// 			expectedOutput: models.TokenUsers{
// 				Users: models.UserDeatilsResponse{
// 					Id:    2,
// 					Name:  "gokul",
// 					Email: "gokul@gmail.com",
// 					Phone: "6282288363",
// 				},
// 				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
// 			},
// 			expectedError: nil,
// 		},
// 	}
// 	for _, test := range testData {
// 		test.StubDetails(*userRepo, *helper, test.input)

// 		tokenusers, err := userUseCase.UserSignUp(test.input.user, test.input.ref)
// 		assert.Equal(t, test.expectedOutput.Users.Id, tokenusers.Users.Id)
// 		assert.Equal(t, test.expectedOutput.Users.Name, tokenusers.Users.Name)
// 		assert.Equal(t, test.expectedOutput.Users.Email, tokenusers.Users.Email)
// 		assert.Equal(t, test.expectedError, err)

// 	}

// }
