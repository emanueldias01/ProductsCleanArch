package router

import "github.com/emanueldias01/ProductsCleanArch/controller"

type ProductRouter struct {
	controller controller.ProductController
}

func NewProductRouter(c controller.ProductController) ProductRouter{
	return ProductRouter{
		controller: c,
	}
}