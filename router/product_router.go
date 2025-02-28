package router

import (
	"github.com/emanueldias01/ProductsCleanArch/controller"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	controller controller.ProductController
}

func NewProductRouter(c controller.ProductController) ProductRouter{
	return ProductRouter{
		controller: c,
	}
}

func (r *ProductRouter) AllRoutes(){
	server := gin.Default()
	server.GET("/products", r.controller.GetProducts)

	defer server.Run()
}