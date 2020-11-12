package database

import (
	"encoding/json"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
	"io/ioutil"
	"os"
)

type Database struct {
	Products 	[]types.Product 	`json:"products"`
	Recipes 	[]types.Recipe 		`json:"recipes"`
}

func (db *Database) GetProduct(barcode string) types.Product {
	for _, product := range db.Products {
		if product.Barcode == barcode {
			return product
		}
	}
	return types.Product{
		Barcode: "",
	}
}

func (db *Database) GetRecipe(id int) types.Recipe {
	for _, recipe := range db.Recipes {
		if recipe.ID == id {
			return recipe
		}
	}
	return types.Recipe{
		ID: -1,
	}
}

func (db *Database) Open() error {
	f, err := os.Open("../../data.json")
	defer f.Close()
	if err != nil {
		return err
	}
	byteValue, _ := ioutil.ReadAll(f)
	json.Unmarshal(byteValue, db)

	return nil
}

func (db *Database) Save() error {
	file, err := json.MarshalIndent(*db, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("../../data.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}