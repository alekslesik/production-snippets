package data

import "gorm.io/gorm"

// Create a Models struct which wraps the MovieModel and UserModel
type Models struct {
	Products   ProductModel
	Categories CategoryModel
	Currencies CurrencyModel
	Images     ImageModel
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel and UserModel
func NewModels(db *gorm.DB) Models {
	return Models{
		Products:   ProductModel{DB: db},
		Categories: CategoryModel{DB: db},
		Currencies: CurrencyModel{DB: db},
		Images:     ImageModel{DB: db},
	}
}
