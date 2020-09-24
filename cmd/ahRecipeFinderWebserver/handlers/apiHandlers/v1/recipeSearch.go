package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetRecipeSearch(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var req struct{
		Ingredients []string `json:"ingredients"`
	}
	json.Unmarshal(reqBody, &req)
	json.NewEncoder(w).Encode(req)
}
