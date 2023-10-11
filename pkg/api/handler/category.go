package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/domain"
	"gokul.go/pkg/usecase/usecaseInterfaces"
	"gokul.go/pkg/utils/models"
	"gokul.go/pkg/utils/response"
)

type CategoryHandler struct {
	categoryUseCase usecaseInterfaces.CategoryUseCase
}

func NewCategoryHandler(usecase usecaseInterfaces.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: usecase}
}
func (cat *CategoryHandler) AddCategory(c *gin.Context) {

	var category domain.Category

	if err := c.BindJSON(&category); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	CategoryResponse, err := cat.categoryUseCase.AddCategory(category)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added category", CategoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}
func (cat *CategoryHandler) UpdateCategory(c *gin.Context) {

	var p models.SetNewName

	if err := c.BindJSON(&p); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilda provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	id := c.Param("id")
	a, err := cat.categoryUseCase.UpdateCategory(id, p.Category)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not update the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully renamed the category", a, nil)
	c.JSON(http.StatusOK, successRes)
}
func (cat *CategoryHandler) DeleteCategory(c *gin.Context) {

	categoryID := c.Query("id")

	fmt.Println("categoryID is", categoryID)

	err := cat.categoryUseCase.DeleteCategory(categoryID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "feilds provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Sucessfully DELETED the category", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
