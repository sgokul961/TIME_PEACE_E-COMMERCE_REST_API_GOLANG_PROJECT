package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"gokul.go/pkg/domain"
	helper_interface "gokul.go/pkg/helper/interface"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	AdminRepository interfaces.AdminRepository
	helper          helper_interface.Helper
}

func NewAdminUseCase(repo interfaces.AdminRepository, help helper_interface.Helper) usecaseInterfaces.AdminUseCase {
	return &adminUseCase{
		AdminRepository: repo, helper: help,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	//getting details of the admin based on the emaim provided

	adminCompareDetails, err := ad.AdminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	//compare password that provided from the databse provided from admins

	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	fmt.Println(err)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	//copy all details except password and sent it back to the front end

	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)

	if err != nil {
		return domain.TokenAdmin{}, err
	}
	tokenString, err := ad.helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil

}
func (ad *adminUseCase) BlockUser(id string) error {

	user, err := ad.AdminRepository.GetUserById(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("already blocked")

	} else {
		user.Blocked = true
	}
	err = ad.AdminRepository.UpdateBlockUserID(user)
	if err != nil {
		return err
	}
	return nil
}

// unblock user
func (ad *adminUseCase) UnblockUser(id string) error {

	user, err := ad.AdminRepository.GetUserById(id)

	if err != nil {
		return err
	}
	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")

	}
	err = ad.AdminRepository.UpdateBlockUserID(user)
	if err != nil {
		return err
	}
	return nil

}
func (ad *adminUseCase) GetUsers(page int, count int) ([]models.UserDetailsAdmin, error) {

	userDetails, err := ad.AdminRepository.GetUsers(page, count)
	if err != nil {
		return []models.UserDetailsAdmin{}, err
	}
	return userDetails, nil
}
