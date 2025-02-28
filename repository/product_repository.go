package repository

import "gorm.io/gorm"

type ProductRepository struct {
	conn *gorm.DB
}

func NewProductRepository(c *gorm.DB) ProductRepository{
	return ProductRepository{
		conn: c,
	}
}