package controller

import "github.com/emanueldias01/ProductsCleanArch/usecase"

type ProductController struct {
	usecase usecase.ProductUseCase
}

func NewProductController(u usecase.ProductUseCase) ProductController{
	return ProductController{
		usecase: u,
	}
}