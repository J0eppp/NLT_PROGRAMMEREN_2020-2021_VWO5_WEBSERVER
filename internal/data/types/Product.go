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
	// Product.SetMainProductCategories splits the Product.MainCategory on every ", " and saves the []string to Product.MainCategories
	// :returns nil
	product.MainCategories = strings.Split(strings.ToLower(product.MainCategory), ", ")
}