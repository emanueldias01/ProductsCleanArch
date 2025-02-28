package model

type Product struct {
	ID int `json:"id"`
	Name string `json:"name_product"`
	Quantity int `json:"quantity_stored`
	Price float64 `json:"price"`
}