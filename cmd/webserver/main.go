package main

import (
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/cmd/webserver/handlers"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/cmd/webserver/handlers/apiHandlers"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/pkg/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.Logger)

	router.HandleFunc("/", handlers.Home)

	apiRouter.HandleFunc("/", apiHandlers.Api)

	log.Fatal(http.ListenAndServe(":8000", router))
}