package controller

import (
	"net/http"

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