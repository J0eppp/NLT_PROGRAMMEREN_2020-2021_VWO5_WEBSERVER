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

	// Get the product name / barcode
	recipeID, err := strconv.ParseInt(_recipeID[0], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	var recipe types.Recipe

	rows, err := memory.DB.Query("SELECT * FROM recipes WHERE ID = ?", recipeID)
	defer rows.Close()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	rows.Next()
	rows.Scan(&recipe.ID, &recipe.URL, &recipe.ImageURL, &recipe.Name)
	if recipe.ID == 0 {
		fmt.Fprintf(w, "{ \"error\": true, \"message\": \"There is no recipe with that ID\" }")
		return
	}

	rows, err = memory.DB.Query("SELECT name FROM ingredients WHERE recipeID = ?", recipe.ID)
	defer rows.Close()
	for rows.Next() {
		var ingredient types.Ingredient
		rows.Scan(&ingredient.Ingredient)
		recipe.Ingredients = append(recipe.Ingredients, ingredient.Ingredient)
	}

	json.NewEncoder(w).Encode(recipe)
}
