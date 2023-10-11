package usecaseInterfaces

import "gokul.go/pkg/domain"

type CategoryUseCase interface {
	AddCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(id string, category string) (domain.Category, error)
	DeleteCategory(CategoryID string) error
}
