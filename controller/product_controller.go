package controller

import (
	"net/http"
	"strconv"

	"github.com/emanueldias01/ProductsCleanArch/model"
	"github.com/emanueldias01/ProductsCleanArch/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	usecase usecase.ProductUseCase
}

func NewProductController(u usecase.ProductUseCase) ProductController{
	return ProductController{
		usecase: u,
	}
}

func (c *ProductController) GetProducts(ctx *gin.Context){
	ctx.JSON(http.StatusOK, c.usecase.GetProducts())
}

func (c *ProductController) CreateProduct(ctx *gin.Context){
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message" : err.Error(),
		})
		return
	}

	model, err := c.usecase.CreateProduct(product)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model)
}

func (c *ProductController) GetProductById(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message" : err.Error(),
		})
		return
	}

	product, err := c.usecase.GetProductById(id)

	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Params.ByName("id"))

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message" : err.Error(),
		})

		return
	}

	c.usecase.DeleteProduct(id)

	ctx.JSON(http.StatusNoContent, nil)
}