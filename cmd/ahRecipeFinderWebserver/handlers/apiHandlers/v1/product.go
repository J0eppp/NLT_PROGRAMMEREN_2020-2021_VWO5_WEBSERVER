package v1

import (
	"encoding/json"
	"fmt"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/memory"
	"net/http"
	"strconv"
	"unicode"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// handlers.apiHandlers.GetProduct gets a product from the AH API and sends it back to the user
	// :param w: http.ResponseWriter
	// :param r: *http.Request
	// :return nil

	// Get the URL param 'product'
	product, ok := r.URL.Query()["product"]
	if !ok || len(product[0]) < 1 {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ 'error': true }")
		return
	}

	// Get the product name / barcode
	var productName = product[0]

	// Check if the entered productName is numeric (barcode) or a string (name)
	var isNumeric = true
	for _, c := range productName {
		if !unicode.IsDigit(c) {
			isNumeric = false
		}
	}

	var p types.Product
	var err error

	if isNumeric {
		// The entered productName is a barcode
		// Check if the product is already saved in the DB
		//rows, err := memory.DB.Query("SELECT barcode, title, mainCategory, subCategory, brand FROM products WHERE barcode = ?", productName)
		barcode, _ := strconv.Atoi(productName)
		p = memory.DB.GetProduct(barcode)
		if p.Barcode == -1 {
			// No result was found..
			p, err = memory.AHConnector.GetProductByBarcode(productName)
			if err != nil {
				fmt.Fprintf(w, "{ 'error': true, 'message': '%s' }", err)
				return
			}
			barcode, err := strconv.ParseInt(productName, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "{ 'error': true, 'message': '%s' }", err)
				return
			}
			p.Barcode = int(barcode)

			//err = p.SaveToDatabase(memory.DB)
			memory.DB.Products = append(memory.DB.Products, p)
			memory.DB.Save()
			if err != nil {
				json.NewEncoder(w).Encode(err)
				return
			}
		} else {
			// Set the main categories
			p.SetMainProductCategories()
			//	rows, err := memory.DB.Query("SELECT `width`, `height`, `URL` FROM `images` WHERE `barcode` = ?", productName)
			//	defer rows.Close()
			//	if err != nil {
			//		fmt.Fprintf(w, "{ 'error': true, 'message': '%s' }", err)
			//		return
			//	}
			//	for rows.Next() {
			//		var img types.Image
			//		rows.Scan(&img.Width, &img.Height, &img.URL)
			//		p.Images = append(p.Images, img)
			//	}
			//}
			//} else {
			// The entered productName is a name (string)
			p, err = memory.AHConnector.GetProductByQuery(productName)
			if err != nil {
				fmt.Fprintf(w, "{ 'error': true, 'message': '%s' }", err)
				return
			}
		}
	}

	// Return the product object to the user 
	json.NewEncoder(w).Encode(p)
}