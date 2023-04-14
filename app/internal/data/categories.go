package data

import (
	"gorm.io/gorm"
)

type Category struct {
}

type CategoryModel struct {
	DB *gorm.DB
}

func (c *CategoryModel) GetAll() ([]*Category, error) {
	return nil, nil
}

func (c *CategoryModel) Insert(product *Category) error {
	return nil
}

func (c CategoryModel) Get(id int64) (*Category, error) {
	return nil, nil
}

func (c CategoryModel) Update(product *Category) error {
	return nil
}

func (c CategoryModel) Delete(id int64) error {
	return nil
}
