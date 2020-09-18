package v1

import (
	"fmt"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"message\": \"Hello there!\" }")
}