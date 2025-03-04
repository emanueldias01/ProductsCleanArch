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
	server.POST("/product", r.controller.CreateProduct)
	server.GET("/product/:id", r.controller.GetProductById)
	server.DELETE("/product/:id", r.controller.DeleteProduct)
	server.PUT("/product", r.controller.UpdateProduct)
	server.POST("/sell/:id", r.controller.SellProduct)
	server.POST("sell/:id/:quantity", r.controller.SellOnScale)

	defer server.Run()
}