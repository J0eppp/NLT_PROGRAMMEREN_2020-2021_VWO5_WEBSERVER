package apiHandlers

import (
	"encoding/json"
	"fmt"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/memory"
	"net/http"
	"unicode"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// handlers.apiHandlers.GetProduct gets a product from the AH API and sends it back to the user
	// :param w: http.ResponseWriter
	// :param r: *http.Request
	// :returns nil

	// Get the URL param 'product'
	product, ok := r.URL.Query()["product"]
	if !ok || len(product[0]) < 1 {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ 'error': true }")
		return
	}

	// Get the product name / barcode
	var productName = product[0]

	fmt.Println("Product name: ", productName)

	// Check if the entered productName is numeric (barcode) or a string (name)
	var isNumeric = true
	for _, c := range productName {
		if !unicode.IsDigit(c) {
			isNumeric = false
		}
	}

	if isNumeric {
		// The entered productName is a barcode

		p, err := memory.AHConnector.GetProductByBarcode(productName)
		if err != nil {
			fmt.Fprintf(w, "{ 'error': true, 'message': '%s' }", err)
			return
		}

		json.NewEncoder(w).Encode(p)
		return
	}


	fmt.Fprintf(w, "{ 'error': false }")
}
