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

func(r *ProductRepository) GetProducts() []model.Product{
	var products []model.Product
	r.conn.Find(&products)
	return products
}