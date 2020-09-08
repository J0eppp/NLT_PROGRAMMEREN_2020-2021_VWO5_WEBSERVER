package main

import (
	"fmt"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
)

func main() {
	var ah = types.AHConnector{}
	ah.GetAnonymousAccessToken()

	ah.GetProductByBarcode("8718906445338")
	var product = ah.GetProductByBarcode("8717677334117")
	fmt.Printf("%+v\n", product)
}
