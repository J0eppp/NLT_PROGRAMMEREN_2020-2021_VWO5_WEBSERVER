package types

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AHConnector struct {
	AccessToken string `json:"accessToken"`
}

func (ah *AHConnector) GetProductByBarcode(barcode string) Product {
	var p = Product{}

	req, err := http.NewRequest("GET", "https://ms.ah.nl/mobile-services/product/search/v1/gtin/" + barcode, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", "Bearer " + ah.AccessToken)
	req.Header.Set("Host", "ms.ah.nl")
	req.Header.Set("User-Agent", "android/6.29.3 Model/phone Android/7.0-API24")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &p)

	p.SetMainProductCategories()

	return p
}

func (ah *AHConnector) GetAnonymousAccessToken() {
	requestBody, err := json.Marshal(map[string]string{
		"client": "appie-anonymous",
	})
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://ms.ah.nl/create-anonymous-member-token", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "ms.ah.nl")
	req.Header.Set("User-Agent", "android/6.29.3 Model/phone Android/7.0-API24")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var r struct{
		AccessToken string `json:"access_token"`
	}

	json.Unmarshal(body, &r)

	ah.AccessToken = r.AccessToken
}