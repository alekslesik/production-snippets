package data

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string  `json:"name"`
	Descrition    *string `json:"descrition"`
	Price         int     `json:"price"`
	CurrencyId    int     `json:"currency_id"`
	Rating        *int    `json:"rating"`
	CategoryId    string  `json:"category_id"`
	Specification *string `json:"specification"`
	ImageId       *string `json:"image_id"`
}

type ProductModel struct {
	*gorm.DB
}

// Get all records from public.products table
func (p *ProductModel) GetAll() ([]Product, error) {
	allProducts := make([]Product, 0)

	// SELECT * FROM products;
	err := p.Find(&allProducts).Error
	if err != nil {
		return nil, err
	}

	return allProducts, nil
}

// Get first object from public.products sorting by primary key ASC
func (p *ProductModel) GetFirst() (Product, error) {
	var product Product

	err := p.First(&product).Error
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p *ProductModel) Insert(product *Category) error {
	return nil
}

func (p ProductModel) Get(id int64) (*Category, error) {
	return nil, nil
}

func (p ProductModel) Update(product *Category) error {
	return nil
}

func (p ProductModel) Delete(id int64) error {
	return nil
}
