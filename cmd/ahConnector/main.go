package main

import (
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
)

func main() {
	var ah = types.AHConnector{}
	ah.GetAnonymousAccessToken()

	//ah.GetProductByBarcode("8718906445338")
	// product, err := ah.GetProductByBarcode("8717677334117")
	// ah.GetProductByQuery("cola")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", product)

	ah.GetSubCategories(1555)
}