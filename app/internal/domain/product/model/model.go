package model

import "time"

type Prooduct struct {
	ID            string    `json:"ID"`
	Name          string    `json:"name"`
	Descrition    *string   `json:"descrition"`
	Price         int       `json:"price"`
	CurrencyId    int       `json:"currency_id"`
	Rating        *int      `json:"rating"`
	CategoryId    string    `json:"category_id"`
	Specification *string   `json:"specification"`
	ImageId       *string   `json:"image_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
