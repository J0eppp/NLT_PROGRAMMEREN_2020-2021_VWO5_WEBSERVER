package v1

import (
	"encoding/json"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/memory"
	"io/ioutil"
	"net/http"
)

func GetRecipeSearch(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var req struct{
		Ingredients []string `json:"ingredients"`
	}
	json.Unmarshal(reqBody, &req)

	var ingredients []types.Ingredient

	// Search for the ingredients in the database
	for _, ingredient := range req.Ingredients {
		rows, err := memory.DB.Query("SELECT * FROM ingredients WHERE name LIKE ?", "%" + ingredient + "%")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		var i types.Ingredient
		rows.Next()
		rows.Scan(&i.ID, &i.Ingredient, &i.RecipeID)
		if i.ID != 0 {
			ingredients = append(ingredients, i)
		}
	}

	json.NewEncoder(w).Encode(ingredients)
}
