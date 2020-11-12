package v1

import (
	"encoding/json"
	"fmt"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/memory"
	"net/http"
	"strconv"
)

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	// Get the URL param 'recipe'
	_recipeID, ok := r.URL.Query()["recipe"]
	if !ok || len(_recipeID[0]) < 1 {
		// Return HTTP error to the user
		fmt.Fprintf(w, "{ 'error': true }")
		return
	}

	// Get the recipe ID
	recipeID, err := strconv.ParseInt(_recipeID[0], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	var recipe types.Recipe

	// Get the recipe from the database based on the ID
	recipe = memory.DB.GetRecipe(int(recipeID))
	if recipe.ID == -1 {
		fmt.Fprintf(w, "{ \"error\": true, \"message\": \"There is no recipe with that ID\" }")
		return
	}

	// Send the recipe object to the user
	json.NewEncoder(w).Encode(recipe)
}
