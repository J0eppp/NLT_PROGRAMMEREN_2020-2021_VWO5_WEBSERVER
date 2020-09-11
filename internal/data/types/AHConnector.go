package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AHConnector struct {
	AccessToken string `json:"accessToken"`
}

func (ah *AHConnector) GetCategories() {
	// AHConnector.GetCategories returns all the categories available
	// :return nil

	req, err := http.NewRequest("GET", "https://ms.ah.nl/mobile-services/v1/product-shelves/categories", nil)

	if err != nil {
		return
	}

	// Set the right headers
	req.Header.Set("Authorization", "Bearer " + ah.AccessToken)
	req.Header.Set("Host", "ms.ah.nl")
	req.Header.Set("User-Agent", "android/6.29.3 Model/phone Android/7.0-API24")

	// Execute the request object
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// Read the response into a []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))
}

func (ah *AHConnector) GetSubCategories(categoryId int) {
	// AHConnector.GetSubCategories returns all the available subcategories
	// :param categoryId: int

	// Create a request object
	req, err := http.NewRequest("GET", "https://ms.ah.nl/mobile-services/product/detail/v3/fir/" + string(categoryId), nil)
	if err != nil {
		return
	}

	// Set the right headers
	req.Header.Set("Authorization", "Bearer " + ah.AccessToken)
	req.Header.Set("Host", "ms.ah.nl")
	req.Header.Set("User-Agent", "android/6.29.3 Model/phone Android/7.0-API24")

	// Execute the request object
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// Read the response into a []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))
}

func (ah *AHConnector) GetProductByBarcode(barcode string) (Product, error) {
	// AHConnector.GetProductByBarcode gets the product information according to the barcode
	// :param barcode: string
	// :return Product
	// :return error

	fmt.Println("Barcode..")

	var p = Product{}

	// Create a request object
	req, err := http.NewRequest("GET", "https://ms.ah.nl/mobile-services/product/search/v1/gtin/" + barcode, nil)
	if err != nil {
		return p, err
	}

	// Set the right headers
	req.Header.Set("Authorization", "Bearer " + ah.AccessToken)
	req.Header.Set("Host", "ms.ah.nl")
	req.Header.Set("User-Agent", "android/6.29.3 Model/phone Android/7.0-API24")

	// Execute the request object
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}

	defer resp.Body.Close()

	// Read the response into a []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return p, err
	}

	// Unmarshal the []byte into *Product
	json.Unmarshal(body, &p)

	// Set the main categories attribute
	p.SetMainProductCategories()

	// Return the product, no error
	return p, nil
}

func (ah *AHConnector) GetProductByQuery(query string) (Product, error) {
	// AHConnector.GetProductByBarcode gets the product information according to the query
	// :param query: string
	// :return Product
	// :return error

	// NOTE: not done yet, response from AH is really weird

	var p = Product{}

	// Set the request body
	requestBody, err := json.Marshal(map[string]string{
		"sortOn": "RELEVANCE",
		"page": "0",
		"size": "0",
		"query": query,
	})
	if err != nil {
		return p, err
	}

	// Create a request object
	req, err := http.NewRequest("GET", "https://ms.ah.nl/mobile-services/product/search/v2?sortOn=RELEVANCE", bytes.NewBuffer(requestBody))
	if err != nil {
		return p, err
	}

	// Set the right headers
	req.Header.Set("Authorization", "Bearer " + ah.AccessToken)
	req.Header.Set("Host", "ms.ah.nl")
	req.Header.Set("User-Agent", "android/6.29.3 Model/phone Android/7.0-API24")

	// Execute the request object
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}

	defer resp.Body.Close()

	// Read the response into a []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return p, err
	}

	fmt.Println(string(body))

	//// Unmarshal the []byte into *Product
	//json.Unmarshal(body, &p)
	//
	//// Set the main categories attribute
	//p.SetMainProductCategories()

	// Return the product, no error
	return p, nil
}

func (ah *AHConnector) GetAnonymousAccessToken() error {
	// AHConnector.GetAnonymousAccessToken requests an authentication token from the AH API and sets AHConnector.AccessToken
	// :return error

	// Create the HTTP POST request body
	requestBody, err := json.Marshal(map[string]string{
		"client": "appie-anonymous",
	})
	if err != nil {
		return err
	}

	// Create a request object
	req, err := http.NewRequest("POST", "https://ms.ah.nl/create-anonymous-member-token", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set the right headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "ms.ah.nl")
	req.Header.Set("User-Agent", "android/6.29.3 Model/phone Android/7.0-API24")

	// Execute the request object
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response into a []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Create a temp struct to store the access token
	var r struct{
		AccessToken string `json:"access_token"`
	}

	// Unmarshal the access token
	json.Unmarshal(body, &r)

	// Set the access token of the AHConnector instance
	ah.AccessToken = r.AccessToken

	// Return nil because there was no error
	return nil
}
