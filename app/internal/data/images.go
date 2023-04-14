package data

import (
	"gorm.io/gorm"
)

type Image struct {
}

type ImageModel struct {
	DB *gorm.DB
}

func (i *ImageModel) GetAll() ([]*Image, error) {
	return nil, nil
}

func (i *ImageModel) Insert(product *Image) error {
	return nil
}

func (i ImageModel) Get(id int64) (*Image, error) {
	return nil, nil
}

func (i ImageModel) Update(product *Image) error {
	return nil
}

func (i ImageModel) Delete(id int64) error {
	return nil
}
