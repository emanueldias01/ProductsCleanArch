package main

import (
	"github.com/emanueldias01/ProductsCleanArch/controller"
	"github.com/emanueldias01/ProductsCleanArch/db"
	"github.com/emanueldias01/ProductsCleanArch/repository"
	"github.com/emanueldias01/ProductsCleanArch/router"
	"github.com/emanueldias01/ProductsCleanArch/usecase"
)

func main() {
	con := db.ConnectDB()
	repository := repository.NewProductRepository(con)
	usecase := usecase.NewProductUseCase(repository)
	controller := controller.NewProductController(usecase)
	server := router.NewProductRouter(controller)

	server.AllRoutes()
}