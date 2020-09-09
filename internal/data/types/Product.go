package types

import "strings"

type Product struct {
	Title string `json:"title"`
	MainCategory string `json:"mainCategory"`
	MainCategories []string `json:"mainCategories"`
	SubCategory string `json:"subCategory"`
	Brand string `json:"brand"`
	Images []Image `json:"images"`
}

func (product *Product) SetMainProductCategories() {
	product.MainCategories = strings.Split(strings.ToLower(product.MainCategory), ", ")
}