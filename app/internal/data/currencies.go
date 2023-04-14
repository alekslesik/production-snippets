package data

import (
	"gorm.io/gorm"
)

type Currency struct {
}

type CurrencyModel struct {
	DB *gorm.DB
}

func (c *CurrencyModel) GetAll() ([]*Currency, error) {
	return nil, nil
}

func (c *CurrencyModel) Insert(product *Currency) error {
	return nil
}

func (c CurrencyModel) Get(id int64) (*Currency, error) {
	return nil, nil
}

func (c CurrencyModel) Update(product *Currency) error {
	return nil
}

func (c CurrencyModel) Delete(id int64) error {
	return nil
}
