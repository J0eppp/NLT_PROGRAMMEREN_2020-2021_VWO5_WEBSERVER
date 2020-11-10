package types

type Recipe struct {
	ID 			int 		`json:"id"`
	Name 		string 		`json:"name"`
	Ingredients []string 	`json:"ingredients"`
	//Tags 		[]string 	`json:"tags"`
	URL 		string 		`json:"url"`
	ImageURL 	string 		`json:"imageURL"`
}
