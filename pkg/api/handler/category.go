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

// @Summary		Add Category
// @Description	Admin can add new categories for products
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			category	body	domain.Category	true	"category"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category/add [post]
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

// @Summary		Update Category
// @Description	Admin can update name of a category into new name
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			set_new_name	body	models.SetNewName	true	"set new name"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category/update [put]
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

// @Summary		Delete Category
// @Description	Admin can delete a category
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category/delete [delete]
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

// ---------------------------------get categories------------------------------------------------//
func (cat *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := cat.categoryUseCase.GetCategories()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot get categories ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Sucessfully got all the categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}
