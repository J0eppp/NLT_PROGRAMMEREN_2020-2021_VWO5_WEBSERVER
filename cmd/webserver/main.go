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
	apiRouter.Use(middleware.SetResponseTypeJSON)

	router.HandleFunc("/", handlers.Home)

	apiRouter.HandleFunc("/", apiHandlers.Api)
	apiRouter.HandleFunc("/product", apiHandlers.GetProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}