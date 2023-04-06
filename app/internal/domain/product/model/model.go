package model

type Prooduct struct {
	ID            string `json:"ID"`
	Name          string `json:"name"`
	Descrition    string `json:"descrition"`
	Price         string `json:"price"`
	CurrencyId    int    `json:"currency_id"`
	Rating        int `json:"rating"`
	CategoryId    string `json:"category_id"`
	Specification string `json:"specification"`
	ImageId       string `json:"image_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
