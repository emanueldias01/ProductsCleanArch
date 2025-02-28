package repository

import (
	"github.com/emanueldias01/ProductsCleanArch/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	conn *gorm.DB
}

func NewProductRepository(c *gorm.DB) ProductRepository{
	return ProductRepository{
		conn: c,
	}
}

func (r *ProductRepository) GetProducts() []model.Product{
	var products []model.Product
	r.conn.Find(&products)
	return products
}

func (r *ProductRepository) CreateProduct(product model.Product) model.Product{
	r.conn.Save(&product)
	return product
}

func (r *ProductRepository) FindProductById(id int) model.Product{
	var product model.Product
	r.conn.First(&product, id)
	return product
}

func (r *ProductRepository) DeleteProduct(id int){
	var product model.Product
	r.conn.Delete(&product, id)
}