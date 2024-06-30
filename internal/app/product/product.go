package product

import (
	"encoding/json"
)

type Product struct {
	ASIN          string `json:"asin"`
	Name          string `json:"name"`
	Brand         string `json:"brand"`
	Price         string `json:"price"`
	PreviousPrice string `json:"previous_price"`
	ChangeDate    string `json:"change_date"`
}

func (p *Product) FormatProduct(asin, inputJSON string) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(inputJSON), &data)
	if err != nil {
		return
	}

	// Assign values to Product struct fields
	p.ASIN = asin
	name, _ := data["name"].(string)
	p.Name = name
	brand, _ := data["brand"].(string)
	p.Brand = brand
	price, _ := data["pricing"].(string)
	p.Price = price
}
