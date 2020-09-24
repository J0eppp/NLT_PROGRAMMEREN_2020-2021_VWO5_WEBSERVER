package main

import (
	"database/sql"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/cmd/ahRecipeFinderWebserver/handlers"
	"github.com/J0eppp/NLT_PROGRAMMEREN_2020-2021_VWO5_WEBSERVER/cmd/ahRecipeFinderWebserver/handlers/apiHandlers/v1"
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
	err := memory.AHConnector.GetAnonymousAccessToken()
	if err != nil {
		panic(err.Error())
	}

	d, err := sql.Open("mysql",  "ahRecipeFinder:Test123@unix(/var/run/mysqld/mysqld.sock)/ahRecipeFinder")
<<<<<<< HEAD
	//d, err := sql.Open("sqlite3", "file:locked.sqlite3?cache=shared")
	//d, err := sql.Open("sqlite3", "./database.db?cache=shared")
=======
>>>>>>> 5670232461047b18d90d7bec35efb77475f071c8
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

	apiV1Router.HandleFunc("/", v1.Api)
	apiV1Router.HandleFunc("/product", v1.GetProduct).Methods("GET")
	apiV1Router.HandleFunc("/product/search", v1.GetProductSearch).Methods("GET")
	apiV1Router.HandleFunc("/recipe/search", v1.GetRecipeSearch).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
