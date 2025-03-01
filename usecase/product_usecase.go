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

	if err != nil{
		return model.Product{}, err
	}

	return u.repository.CreateProduct(product), err
}

func (u *ProductUseCase) GetProductById(id int) (model.Product, error){
	product := u.repository.FindProductById(id)
	var err error
	if product.Name == ""{
		err = errors.New("product not found")
	}

	return product, err
}

func (u *ProductUseCase) UpdateProduct(product model.Product) (model.Product, error) {
	productDB, err := u.GetProductById(product.ID)
	if err != nil {
		return model.Product{}, err
	}

	if product.Price < 0 {
		return model.Product{}, errors.New("invalid price")
	}
	if product.Quantity < 0 {
		return model.Product{}, errors.New("invalid quantity")
	}

	if product.Name != "" {
		productDB.Name = product.Name
	}
	if product.Price != 0 {
		productDB.Price = product.Price
	}
	if product.Quantity != 0 {
		productDB.Quantity = product.Quantity
	}

	return u.repository.UpdateProduct(productDB)
}


func (u *ProductUseCase) DeleteProduct(id int){
	u.repository.DeleteProduct(id)
}

func (u *ProductUseCase) SellProduct(id int) (model.Product, error){
	product, err := u.GetProductById(id)

	if err != nil{
		return model.Product{}, err
	}

	if product.Quantity == 1{
		u.DeleteProduct(id)
		return product, nil
	}

	product.Quantity--
	return u.UpdateProduct(product)
}

func (u *ProductUseCase) SellOnScale(id int, quantity int) (model.Product, error){
	product, err := u.GetProductById(id)

	if err != nil{
		return model.Product{}, err
	}

	if product.Quantity < quantity{
		return model.Product{}, errors.New("not enough products")
	}

	if product.Quantity == quantity{
		u.DeleteProduct(id)
		return product, nil
	}

	product.Quantity -= quantity
	u.UpdateProduct(product)
	return product, nil
}