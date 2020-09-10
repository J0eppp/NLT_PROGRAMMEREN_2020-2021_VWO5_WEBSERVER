package types

import (
	"database/sql"
	"errors"
	"strings"
)

type Product struct {
	Title string `json:"title"`
	Barcode int `json:"barcode"`
	MainCategory string `json:"mainCategory"`
	MainCategories []string `json:"mainCategories"`
	SubCategory string `json:"subCategory"`
	Brand string `json:"brand"`
	Images []Image `json:"images"`
}

func (product *Product) SetMainProductCategories() {
	// Product.SetMainProductCategories splits the Product.MainCategory on every ", " and saves the []string to Product.MainCategories
	// :return nil
	product.MainCategories = strings.Split(strings.ToLower(product.MainCategory), ", ")
}

func (product *Product) SaveToDatabase(db *sql.DB) error {
	// Product.SaveToDatabase saves the product in the database if it does not already exist for caching purposes
	// :param db: *sql.DB
	// :return error

	// Check if the barcode is set properly
	if product.Barcode <= -1 {
		// Barcode is invalid... return error
		return errors.New("invalid product.Barcode")
	}

	// Barcode is valid, check if it doesn't already exist in the database
	rows, err := db.Query("SELECT id FROM products WHERE barcode = ?", product.Barcode)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	var id int
	rows.Scan(&id)
	if id != 0 {
		// Barcode is already saved, return nil because there was no error
		return nil
	}

	// Barcode doesn't exist already, save it to the database
	statement, err := db.Prepare("INSERT INTO products (barcode, title, mainCategory, subCategory, brand) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(product.Barcode, product.Title, product.MainCategory, product.SubCategory, product.Brand)
	if err != nil {
		return err
	}

	// Save the images separately with a reference to this product (by barcode)
	for _, image := range product.Images {
		statement, err = db.Prepare("INSERT INTO images (barcode, width, height, URL) VALUES (?, ?, ?, ?)")
		if err != nil {
			return err
		}

		_, err = statement.Exec(product.Barcode, image.Width, image.Height, image.URL)
		if err != nil {
			return err
		}
	}

	return nil
}