package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaseat/pManager/api"
	"github.com/kaseat/pManager/portfolio"
)

func main() {
	cfg := portfolio.Config{
		MongoURL: "mongodb://localhost:27017",
		DbName:   "tcs",
	}
	portfolio.Init(cfg)

	fmt.Println("Started!")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/portfolios", api.CreateSinglePortfolio).Methods("POST")
	router.HandleFunc("/portfolios", api.ReadAllPortfolios).Methods("GET")
	router.HandleFunc("/portfolios", api.DeleteAllPortfolios).Methods("DELETE")
	router.HandleFunc("/portfolios/{id}", api.ReadSinglePortfolio).Methods("GET")
	router.HandleFunc("/portfolios/{id}", api.UptateSinglePortfolio).Methods("PUT")
	router.HandleFunc("/portfolios/{id}", api.DeleteSinglePortfolio).Methods("DELETE")
	router.HandleFunc("/portfolios/{id}/operations", api.ReadAllOperations).Methods("GET")
	router.HandleFunc("/portfolios/{id}/operations", api.CreateSingleOperation).Methods("POST")
	router.HandleFunc("/portfolios/{id}/operations", api.DeleteAllOperations).Methods("DELETE")
	router.HandleFunc("/portfolios/{id}/operations/{figi}/average", api.GetAveragePrice).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
