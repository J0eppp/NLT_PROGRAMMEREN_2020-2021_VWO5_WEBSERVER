package v1

import (
	"encoding/json"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/memory"
	"io/ioutil"
	"net/http"
	"strings"
)

func like(base string, sub string) bool {
	return strings.Contains(strings.ToLower(base), strings.ToLower(sub))
}

func GetRecipeSearch(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	// Get the ingredients sent by the user
	var req struct{
		Ingredients []string `json:"ingredients"`
	}
	json.Unmarshal(reqBody, &req)

	var ingredients []types.Ingredient

	for _, ingredient := range req.Ingredients {
		for _, recipe := range memory.DB.Recipes {
			for _, i := range recipe.Ingredients {
				if like(i, ingredient) {
					ingredients = append(ingredients, types.Ingredient{
						Ingredient: i,
						RecipeID: recipe.ID,
					})
				}
			}
		}
	}

	// Return all the ingredient objects to the user
	json.NewEncoder(w).Encode(ingredients)
}
