package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gokul.go/pkg/config"
	"gokul.go/pkg/mock"
	"gokul.go/pkg/utils/models"
)

// func Test_User_UserSignUp(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	// Create mock implementations for the repositories

// 	userRepo := mock.NewMockUserRepository(ctrl)
// 	orderRepo := mock.NewMockOrderRepository(ctrl)
// 	otpRepo := mock.NewMockOtpRepository(ctrl)
// 	helper := mock.NewMockHelper(ctrl)
// 	cfg := config.Config{}

// 	type args struct {
// 		user models.UserDetails
// 		ref  string
// 	}

// 	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, helper, orderRepo)
// 	testData := map[string]struct {
// 		input          args
// 		StubDetails    func(mock.MockUserRepository, mock.MockHelper, args)
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

// 			StubDetails: func(userRepo mock.MockUserRepository, helper mock.MockHelper, signupData args) {
// 				gomock.InOrder(
// 					userRepo.EXPECT().CheckUserAvailability(signupData.user.Email).Times(1).Return(false),
// 					userRepo.EXPECT().FindUserFromReference(signupData.ref).Times(1).Return(1, nil),
// 					helper.EXPECT().GenerateRefferalCode().Times(1).Return("goku", nil),
// 					userRepo.EXPECT().UserSignUp(signupData.user, "goku").Times(1).Return(
// 						func(user models.UserDetails, referral string) models.UserDeatilsResponse {
// 							// Simulate password hashing before storing in the database
// 							hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 							return models.UserDeatilsResponse{
// 								Id:    2,
// 								Name:  user.Name,
// 								Email: user.Email,
// 								Phone: user.Phone,
// 							}
// 						},
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

// func Test_User_LoginHandler(t *testing.T) {
// 	cntrl := gomock.NewController(t)
// 	// Create mock implementations for the repositories

// 	userRepo := mocks.NewMockUserRepository(cntrl)
// 	helper := mocks.NewMockHelper(cntrl)
// 	cfg := config.Config{}

// 	type args struct {
// 		user models.UserLoign
// 	}

// 	userUseCase := NewUserUseCase(userRepo, cfg, nil, nil, helper)
// 	testData := map[string]struct {
// 		input          args
// 		StubDetails    func(mocks.MockUserRepository, mocks.MockHelper, args)
// 		expectedOutput models.TokenUsers
// 		expectedError  error
// 	}{
// 		"success": {
// 			input: args{
// 				user: models.UserLoign{
// 					Email:    "gokul@gmail.com",
// 					PassWord: "hffj",
// 				},
// 			},

// 			StubDetails: func(userRepo mocks.MockUserRepository, helper mocks.MockHelper, loginData args) {
// 				gomock.InOrder(
// 					userRepo.EXPECT().CheckUserAvailability(loginData.user.Email).Times(1).Return(true),
// 					userRepo.EXPECT().UserBlockStatus(loginData.user.Email).Times(1).Return(false, nil),
// 					userRepo.EXPECT().FindUserByEmail(loginData.user).Times(1).Return(
// 						models.UserSignInResponse{
// 							ID:       1,
// 							Email:    loginData.user.Email,
// 							Password: "$2a$10$kgGwgsjgQ.ycqVn.BHRwT.FvW57fvV8ncmtYcmWTNMsLvStW4EvQi", // Replace with the hashed password
// 						}, nil),
// 					helper.EXPECT().GenerateTokenClients(models.UserDeatilsResponse{
// 						ID:       1,
// 						Email:    loginData.user.Email,
// 						Password: "$2a$10$kgGwgsjgQ.ycqVn.BHRwT.FvW57fvV8ncmtYcmWTNMsLvStW4EvQi", // Replace with the hashed password
// 					}).Times(1).Return("ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs", nil),
// 				)
// 			},

// 			expectedOutput: models.TokenUsers{
// 				Users: models.UserDeatilsResponse{
// 					ID:       1,
// 					Email:    "gokul@gmail.com",
// 					Password: "", // Password should not be exposed in the response
// 				},
// 				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
// 			},
// 			expectedError: nil,
// 		},
// 		"user not found": {
// 			input: args{
// 				user: models.UserLoign{
// 					Email:    "nonexistent@gmail.com",
// 					PassWord: "hffj",
// 				},
// 			},

// 			StubDetails: func(userRepo mocks.MockUserRepository, helper mocks.MockHelper, loginData args) {
// 				userRepo.EXPECT().CheckUserAvailability(loginData.user.Email).Times(1).Return(false)
// 			},

// 			expectedOutput: models.TokenUsers{},
// 			expectedError:  ErrUserNotFound,
// 		},
// 		"user blocked by admin": {
// 			input: args{
// 				user: models.UserLoign{
// 					Email:    "blocked@gmail.com",
// 					PassWord: "hffj",
// 				},
// 			},

// 			StubDetails: func(userRepo mocks.MockUserRepository, helper mocks.MockHelper, loginData args) {
// 				gomock.InOrder(
// 					userRepo.EXPECT().CheckUserAvailability(loginData.user.Email).Times(1).Return(true),
// 					userRepo.EXPECT().UserBlockStatus(loginData.user.Email).Times(1).Return(true, nil),
// 				)
// 			},

// 			expectedOutput: models.TokenUsers{},
// 			expectedError:  ErrBlockedByAdmin,
// 		},
// 		"incorrect password": {
// 			input: args{
// 				user: models.UserLoign{
// 					Email:    "gokul@gmail.com",
// 					PassWord: "wrongpassword",
// 				},
// 			},

// 			StubDetails: func(userRepo mocks.MockUserRepository, helper mocks.MockHelper, loginData args) {
// 				gomock.InOrder(
// 					userRepo.EXPECT().CheckUserAvailability(loginData.user.Email).Times(1).Return(true),
// 					userRepo.EXPECT().UserBlockStatus(loginData.user.Email).Times(1).Return(false, nil),
// 					userRepo.EXPECT().FindUserByEmail(loginData.user).Times(1).Return(
// 						models.UserSignInResponse{
// 							ID:       1,
// 							Email:    loginData.user.Email,
// 							Password: "$2a$10$kgGwgsjgQ.ycqVn.BHRwT.FvW57fvV8ncmtYcmWTNMsLvStW4EvQi", // Replace with the hashed password
// 						}, nil),
// 					helper.EXPECT().GenerateTokenClients(gomock.Any()).Times(0),
// 				)
// 			},

// 			expectedOutput: models.TokenUsers{},
// 			expectedError:  ErrIncorrectPassword,
// 		},
// 	}

// 	for _, test := range testData {
// 		test.StubDetails(*userRepo, *helper, test.input)

// 		tokenusers, err := userUseCase.LoginHandler(test.input.user)
// 		assert.Equal(t, test.expectedOutput.Users.ID, tokenusers.Users.ID)
// 		assert.Equal(t, test.expectedOutput.Users.Email, tokenusers.Users.Email)
// 		assert.Equal(t, "", tokenusers.Users.Phone) // Ensure the password is not exposed in the response
// 		assert.Equal(t, test.expectedOutput.Token, tokenusers.Token)
// 		assert.Equal(t, test.expectedError, err)
// 	}
// 	cntrl.Finish()
// }

func Test_User_UserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock implementations for the repositories
	userRepo := mock.NewMockUserRepository(ctrl)
	orderRepo := mock.NewMockOrderRepository(ctrl)
	otpRepo := mock.NewMockOtpRepository(ctrl)
	helper := mock.NewMockHelper(ctrl)
	cfg := config.Config{}

	type args struct {
		user models.UserDetails
		ref  string
	}

	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, helper, orderRepo)

	testData := map[string]struct {
		input          args
		StubDetails    func(mock.MockUserRepository, mock.MockHelper, args)
		expectedOutput models.TokenUsers
		expectedError  error
	}{
		"success": {
			input: args{
				user: models.UserDetails{
					Name:            "gokul",
					Email:           "gokul@gmail.com",
					Phone:           "6282288363",
					Password:        "hffj",
					ConfirmPassWord: "hffj",
				},
				ref: "gok",
			},
			StubDetails: func(userRepo mock.MockUserRepository, helper mock.MockHelper, arg args) {
				gomock.InOrder(
					userRepo.EXPECT().CheckUserAvailability(arg.user.Email).Times(1).Return(false),
					userRepo.EXPECT().FindUserFromReference("1234").Times(1).Return(false),
					helper.EXPECT().PasswordHashing(arg.user.Password).Times(1).Return(gomock.Any().String(), nil),
					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
					userRepo.EXPECT().UserSignUp(arg.user, gomock.Any().String()).Times(1).Return(
						models.UserDeatilsResponse{
							Id:    1,
							Name:  arg.user.Name,
							Email: arg.user.Email,
							Phone: arg.user.Phone,
						}, nil,
					),
					helper.EXPECT().GenerateTokenClients(models.UserDeatilsResponse{
						Id:    1,
						Name:  "gokul",
						Email: "gokul@gmail.com",
						Phone: "6282288363",
					}).Times(1).Return(gomock.Any().String(), nil),
				)
			},
			expectedOutput: models.TokenUsers{
				Users: models.UserDeatilsResponse{
					Id:    1,
					Name:  "gokul",
					Email: "gokul@gmail.com",
					Phone: "6282288363",
				},
				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
			},
			expectedError: nil,
		},
		"user already exist":{
			input: models.UserDetails{
				Name: "gokul",
				Email: "gokul@gmail.com",
				Phone: "6282288363",
				Password: gomock.Any().String(),
				ConfirmPassWord: "gokk",
			},
			StubDetails: func(userRepo mock.MockUserRepository, helper mock.MockHelper, arg args) {},
		}
	}

}
