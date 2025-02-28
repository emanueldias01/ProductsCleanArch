package usecase

import (
	"github.com/emanueldias01/ProductsCleanArch/model"
	"github.com/emanueldias01/ProductsCleanArch/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) ProductUseCase{
	return ProductUseCase{
		repository: r,
	}
}

func (u *ProductUseCase) GetProducts() []model.Product{
	return u.repository.GetProducts()
}