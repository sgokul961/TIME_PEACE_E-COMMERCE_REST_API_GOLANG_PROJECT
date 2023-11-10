package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
)

type categoryUseCase struct {
	repository interfaces.CategoryRepository
}

func NewCategoryUseCase(repo interfaces.CategoryRepository) usecaseInterfaces.CategoryUseCase {
	return &categoryUseCase{
		repository: repo,
	}
}
func (cat *categoryUseCase) AddCategory(category domain.Category) (domain.Category, error) {

	if err := cat.repository.CheckCategories(category); err != nil {
		return domain.Category{}, err
	}
	productResponse, err := cat.repository.AddCategory(category)

	if err != nil {
		return domain.Category{}, err
	}
	return productResponse, nil
}

func (cat *categoryUseCase) UpdateCategory(id string, category string) (domain.Category, error) {
	cid, err := strconv.Atoi(id)
	if err != nil {
		return domain.Category{}, err
	}
	result, err := cat.repository.CheckCatogory(uint(cid))

	if err != nil {
		fmt.Println("1")
		return domain.Category{}, err

	}
	if !result {
		fmt.Println("2")
		return domain.Category{}, errors.New("there is no category as you mentioned")

	}
	newcat, err := cat.repository.UpdateCategory(uint(cid), category)
	if err != nil {
		fmt.Println("3")
		return domain.Category{}, err
	}
	return newcat, err

}
func (cat *categoryUseCase) DeleteCategory(categoryID string) error {

	err := cat.repository.DeleteCategory(categoryID)

	if err != nil {
		return err
	}
	return nil
}
func (c *categoryUseCase) GetCategories() ([]domain.Category, error) {
	categories, err := c.repository.GetCategories()
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}
