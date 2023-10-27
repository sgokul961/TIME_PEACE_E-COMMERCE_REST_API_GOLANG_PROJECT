package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"gokul.go/pkg/config"
	"gokul.go/pkg/domain"
	helper_interface "gokul.go/pkg/helper/interface"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo        interfaces.UserRepository
	cfg             config.Config
	otpRepository   interfaces.OtpRepository
	helper          helper_interface.Helper
	orderRepository interfaces.OrderRepository
}

func NewUserUseCase(repo interfaces.UserRepository, cfg config.Config, otp interfaces.OtpRepository, help helper_interface.Helper, order interfaces.OrderRepository) usecaseInterfaces.UserUseCase {
	return &userUseCase{
		userRepo:        repo,
		cfg:             cfg,
		otpRepository:   otp,
		helper:          help,
		orderRepository: order,
	}
}

var InternalError = "Internal Server Error"

func (u *userUseCase) UserSignUp(user models.UserDetails, ref string) (models.TokenUsers, error) {

	fmt.Println("add users")
	//check if user exist
	userExist := u.userRepo.CheckUserAvailability(user.Email)
	fmt.Println("user exist ", userExist)

	if userExist {
		return models.TokenUsers{}, errors.New("user already exist,sign in")
	}
	fmt.Println(user)

	if user.Password != user.ConfirmPassWord {
		return models.TokenUsers{}, errors.New("password dosnt match")
	}
	refernseUser, err := u.userRepo.FindUserFromReference(ref)
	if err != nil {
		return models.TokenUsers{}, errors.New("cannot find reference user")
	}
	//hash password since details are validated

	hashePassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")

	}
	user.Password = string(hashePassword)

	referral, err := u.helper.GenerateRefferalCode()

	if err != nil {
		return models.TokenUsers{}, errors.New(InternalError)
	}

	if err != nil {
		return models.TokenUsers{}, err
	}
	//add user details to the database
	userData, err := u.userRepo.UserSignUp(user, referral)
	if err != nil {
		return models.TokenUsers{}, err
	}

	//create a jwt token string for the user

	tokenString, err := u.helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not crate a token due to some internal error")

	}
	// var userDetails models.UserDeatilsResponse
	// //copies all the details except password from the user.

	// err = copier.Copy(&userDetails, &userData)

	// if err != nil {
	// 	return models.TokenUsers{}, err
	// }

	//credit 20 rupees to the user which is the source of the reference code
	if err := u.userRepo.CreditReferencePointsToWallet(refernseUser); err != nil {
		return models.TokenUsers{}, errors.New("error in crediting gift")
	}
	//creating a new wallet
	if _, err := u.orderRepository.CreateNewWallet(userData.Id); err != nil {
		return models.TokenUsers{}, errors.New("errorin creating wallet")
	}
	return models.TokenUsers{
		Users: userData,
		Token: tokenString,
	}, nil
}
func (u *userUseCase) LoginHandler(user models.UserLoign) (models.TokenUsers, error) {
	//checking if user name exist with this email address.
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)

	if err != nil {
		return models.TokenUsers{}, err
	}

	if isBlocked {
		return models.TokenUsers{}, errors.New("user is already blocked by admin")
	}

	//Get the user details in order to check the password,in this case (the same function can be used for future)

	user_details, err := u.userRepo.FindUserByEmail(user)

	//fmt.Println("user details", user_details)

	if err != nil {
		return models.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.PassWord))

	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDeatilsResponse

	err = copier.Copy(&userDetails, &user_details)
	fmt.Println("user details ", userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := u.helper.GenerateTokenClients(userDetails)

	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")
	}
	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}
func (ia *userUseCase) AddAddress(id int, address models.AddAddress) error {

	err := ia.userRepo.AddAddress(id, address)
	if err != nil {
		return err
	}
	return nil
}
func (i *userUseCase) GetAddress(id int) ([]domain.Address, error) {
	addresses, err := i.userRepo.GetAddress(id)

	if err != nil {
		return []domain.Address{}, err
	}
	return addresses, nil
}
func (i *userUseCase) GetUserDetails(id int) (models.UserDeatilsResponse, error) {

	details, err := i.userRepo.GetUserDetails(id)

	if err != nil {
		return models.UserDeatilsResponse{}, err
	}
	return details, nil
}
func (i *userUseCase) ChangePassword(id int, old string, password string, repassword string) error {

	userPassword, err := i.userRepo.GetPassword(id)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(old))

	if err != nil {
		return errors.New("password incorrect")

	}
	if password != repassword {
		return errors.New("password dose not match")

	}
	newpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New("internal server error")
	}
	return i.userRepo.ChangePassword(id, string(newpassword))
}
func (i *userUseCase) ForgotPasswordSend(phone string) error {

	ok := i.otpRepository.FindUserByMobileNumber(phone)

	if !ok {
		return errors.New("the user does not exist")
	}

	i.helper.TwilioSetup(i.cfg.ACCOUNTSID, i.cfg.AUTHTOKEN)
	fmt.Println("accsid:", i.cfg.ACCOUNTSID)
	fmt.Println("auth:", i.cfg.AUTHTOKEN)

	_, err := i.helper.TwilioSendOTP(phone, i.cfg.SERVICESID)
	if err != nil {
		return errors.New("error occured while genarating OTP")

	}
	return nil

}
func (u *userUseCase) GetCart(id int) ([]models.GetCart, error) {
	//find cart id

	cart_id, err := u.userRepo.GetCartID(id)
	if err != nil {
		return []models.GetCart{}, errors.New("internal error")

	}
	fmt.Println("cart id is:", cart_id)

	//find products inside cart

	products, err := u.userRepo.GetProductsInCart(cart_id)
	if err != nil {
		return []models.GetCart{}, errors.New("internal error")
	}

	fmt.Println("product is :", products)

	//find product names

	var product_names []string

	for i := range products {
		product_name, err := u.userRepo.FindProductNames(products[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error")

		}
		product_names = append(product_names, product_name)
	}
	fmt.Println("product name is :", product_names)
	//find quantity

	var quantity []int

	for i := range products {
		q, err := u.userRepo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error")
		}
		fmt.Println("quantity q:", q)
		quantity = append(quantity, q)
	}
	fmt.Println("the quantity is:", quantity)

	var price []float64

	for i := range products {
		q, err := u.userRepo.FindPrice(products[i])

		if err != nil {
			return []models.GetCart{}, errors.New("internal error")
		}
		price = append(price, q)
	}
	fmt.Println("the price is:", price)
	var categories []int

	for i := range products {
		c, err := u.userRepo.FindCategory(products[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error")

		}
		categories = append(categories, c)
	}
	fmt.Println(categories)
	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart

		get.ProductName = product_names[i]
		get.Category_id = categories[i]
		get.Quantity = quantity[i]
		get.Totel = price[i]
		get.DiscountedPrice = 0

		getcart = append(getcart, get)
	}
	fmt.Println("getcart is:", getcart)
	//find offers
	var offers []int

	for i := range categories {
		c, err := u.userRepo.FindofferPercentage(categories[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error")
		}
		offers = append(offers, c)
	}
	//discounted price
	for i := range getcart {
		getcart[i].DiscountedPrice = (getcart[i].Totel) - (getcart[i].Totel * float64(offers[i]) / 100)
	}
	fmt.Println(getcart)
	//then return in appopriate format

	return getcart, nil

}
func (i *userUseCase) RemoveFromCart(cart_id, inventory_id int) error {

	err := i.userRepo.RemoveFromCart(cart_id, inventory_id)
	if err != nil {
		return err
	}
	return nil
}
func (i *userUseCase) UpdateQuantityAdd(id, inventory_id int) error {

	err := i.userRepo.UpdateQuantityAdd(id, inventory_id)
	if err != nil {
		return err
	}
	return nil
}
func (i *userUseCase) UpdateQuantityMinus(id, inventory_id int) error {

	err := i.userRepo.UpdateQuantityMinus(id, inventory_id)
	if err != nil {
		return err
	}
	return nil
}
func (i *userUseCase) VarifyForgotPasswordAndChange(model models.ForgotVarify) error {

	i.helper.TwilioSetup(i.cfg.ACCOUNTSID, i.cfg.AUTHTOKEN)

	err := i.helper.TwilioVerifyOTP(i.cfg.SERVICESID, model.Otp, model.Phone)

	if err != nil {
		return errors.New("error while varifying ")

	}
	id, err := i.userRepo.FindIdFromPhone(model.Phone)

	if err != nil {
		return errors.New("cannot find user from  mobile number")
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(model.NewPassword), 10)
	if err != nil {
		return errors.New("hashing problem")
	}
	//if user is auhenticated then change the password i the database

	if err := i.userRepo.ChangePassword(id, string(newPassword)); err != nil {
		return err
	}
	return nil

}
func (i *userUseCase) GetMyReferanceLink(id int) (string, error) {
	baseURL := "gokul.go/user/signup"

	referralCode, err := i.userRepo.GetReferralCodeFromID(id)
	if err != nil {
		return "", errors.New("error getting ref code")

	}
	referralLInk := fmt.Sprintf("%s?ref=%s", baseURL, referralCode)

	return referralLInk, nil
}
