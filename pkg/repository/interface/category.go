package interfaces

import "gokul.go/pkg/domain"

type CategoryRepository interface {
	AddCategory(category domain.Category) (domain.Category, error)
	CheckCatogory(id uint) (bool, error)
	UpdateCategory(id uint, category string) (domain.Category, error)
	DeleteCategory(categoryID string) error
}
