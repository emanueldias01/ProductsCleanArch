package usecase

import "github.com/emanueldias01/ProductsCleanArch/repository"

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) ProductUseCase{
	return ProductUseCase{
		repository: r,
	}
}