package usecase

import (
	"errors"

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

func (u *ProductUseCase) CreateProduct(product model.Product) (model.Product, error){
	var err error = nil
	if product.Price < 0{
		err = errors.New("invalid price")
	}
	if product.Name == ""{
		err = errors.New("invalid name")
	}
	if product.Quantity < 0{
		err = errors.New("invalid quantity")
	}

	return u.repository.CreateProduct(product), err
}
