package handler

import (
	"fmt"
	"net/http"
	"toko-belanja/dto"
	"toko-belanja/pkg/errs"
	"toko-belanja/pkg/helpers"
	"toko-belanja/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}

func (p *productHandler) CreateProduct(c *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	res, err := p.productService.CreateProduct(productRequest)

	if err != nil {
		fmt.Println(err)
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (ph *productHandler) GetProducts(ctx *gin.Context) {
	response, err := ph.productService.GetProducts()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (ph *productHandler) UpdateProductById(ctx *gin.Context) {
	var productRequest dto.NewProductRequest
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	productId, err := helpers.GetParamId(ctx, "productId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := ph.productService.UpdateProductById(productId, productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (ph *productHandler) DeleteProduct(ctx *gin.Context) {
	productId, err := helpers.GetParamId(ctx, "productId")
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	result, err := ph.productService.DeleteProduct(productId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}