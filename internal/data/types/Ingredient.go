package types

type Ingredient struct {
	ID 			int 	`json:"ID"`
	Ingredient 	string 	`json:"ingredient"`
	RecipeID 	int 	`json:"recipeID"`
}
