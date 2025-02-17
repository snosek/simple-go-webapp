package data

import (
	"encoding/json"
)

type Product struct {
	Name     string
	Occasion []string
	Size     []string
	Price    map[string]int
}

func GetProducts() ([]Product, error) {
	var products []Product
	data, err := Data.ReadFile("products.json")
	if err != nil {
		return products, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return products, err
	}
	return products, nil
}
