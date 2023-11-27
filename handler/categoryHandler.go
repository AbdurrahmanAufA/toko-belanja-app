package handler

import (
	"net/http"
	"toko-belanja/dto"
	"toko-belanja/entity"
	"toko-belanja/pkg/errs"
	"toko-belanja/pkg/helpers"
	"toko-belanja/service"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) categoryHandler {
	return categoryHandler{
		categoryService: categoryService,
	}
}

// CreateNewCategory godoc
// @Tags categorys
// @Description Create New Category Data
// @ID create-new-category
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewCategoryRequest true "request body json"
// @Success 201 {object} dto.NewCategoryRequest
// @Router /categorys [post]
func (m categoryHandler) CreateNewCategory(c *gin.Context) {
	var categoryRequest dto.NewCategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newCategory, err := m.categoryService.CreateNewCategory(user.Id, categoryRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}

// func (m categoryHandler) UpdateCategoryById(c *gin.Context) {
// 	var categoryRequest dto.NewCategoryRequest

// 	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
// 		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

// 		c.JSON(errBindJson.Status(), errBindJson)
// 		return
// 	}

// 	categoryId, err := helpers.GetParamId(c, "categoryId")

// 	if err != nil {
// 		c.AbortWithStatusJSON(err.Status(), err)
// 		return
// 	}

// 	response, err := m.categoryService.UpdateCategoryById(categoryId, categoryRequest)

// 	if err != nil {
// 		c.AbortWithStatusJSON(err.Status(), err)
// 		return
// 	}

// 	c.JSON(response.StatusCode, response)
// }

func (ch *categoryHandler) GetCategory(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := ch.categoryService.GetCategoryById(user.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (ch *categoryHandler) PatchCategory(ctx *gin.Context) {
	var newCategoryRequest dto.NewCategoryRequest
	if err := ctx.ShouldBindJSON(&newCategoryRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	categoryId, err := helpers.GetParamId(ctx, "categoryId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ch.categoryService.PatchCategory(categoryId, newCategoryRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (ch *categoryHandler) DeleteCategory(ctx *gin.Context) {

	categoryId, err := helpers.GetParamId(ctx, "categoryId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ch.categoryService.DeleteCategory(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
