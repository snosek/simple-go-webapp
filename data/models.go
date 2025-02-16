package models

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Product struct {
	Name     string
	Occasion []string
	Size     []string
	Price    map[string]int
}

func GetProducts() []Product {
	fpath, err := filepath.Abs("data/products.json")
	data, err := os.ReadFile(fpath)
	if err != nil {
		log.Print(err.Error())
	}
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Print(err.Error())
	}
	return products
}
