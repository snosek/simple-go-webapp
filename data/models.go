package data

import (
	"encoding/json"
	"errors"
)

type Product struct {
	Name     string
	Occasion []string
	Size     []string
	Price    map[string]int
}

func getProducts() ([]Product, error) {
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

var Products, _ = getProducts()

func GetProductWithName(name string) (Product, error) {
	for _, product := range Products {
		if product.Name == name {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}
