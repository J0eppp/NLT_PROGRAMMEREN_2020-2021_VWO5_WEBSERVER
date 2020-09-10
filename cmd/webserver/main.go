package main

import (
	"database/sql"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/cmd/webserver/handlers"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/cmd/webserver/handlers/apiHandlers"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/data/types"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/internal/memory"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/pkg/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Setup default things here
	memory.AHConnector = types.AHConnector{}
	memory.AHConnector.GetAnonymousAccessToken()

	d, err := sql.Open("mysql",  "root:Test123@unix(/var/run/mysqld/mysqld.sock)/ahRecipeFinder")
	if err != nil {
		panic(err)
	}

	memory.DB = d

	router := mux.NewRouter().StrictSlash(true)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.SetResponseTypeJSON)

	apiV1Router := apiRouter.PathPrefix("/v1").Subrouter()

	router.HandleFunc("/", handlers.Home)

	apiV1Router.HandleFunc("/", apiHandlers.Api)
	apiV1Router.HandleFunc("/product", apiHandlers.GetProduct).Methods("GET")
	apiV1Router.HandleFunc("/product/search", apiHandlers.GetProductSearch).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}