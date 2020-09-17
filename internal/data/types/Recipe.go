package types

type Recipe struct {
	Name 		string 		`json:"name"`
	Ingredients []string 	`json:"ingredients"`
	Tags 		[]string 	`json:"tags"`
	URL 		[]string 	`json:"url"`
	ImageURL 	string 		`json:"imageURL"`
}
