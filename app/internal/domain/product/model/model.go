package model

type Prooduct struct {
	Name          string `json:"name"`
	Descrition    string `json:"descrition"`
	Price         string `json:"price"`
	CurrencyId    string `json:"currency_id"`
	Rating        string `json:"rating"`
	CategoryId    string `json:"category_id"`
	Specification string `json:"specification"`
	ImageId       string `json:"image_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
