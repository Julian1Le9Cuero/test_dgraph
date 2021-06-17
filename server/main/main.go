package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julian1le9cuero/dgraph_challenge/server/api/buyers"

	// "github.com/julian1le9cuero/dgraph_challenge/server/api/products"
	// "github.com/julian1le9cuero/dgraph_challenge/server/api/transactions"
	"github.com/julian1le9cuero/dgraph_challenge/server/api/utils"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/dgo_test"

	// "github.com/julian1le9cuero/dgraph_challenge/server/database/models"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/storage"
)

func main() {
	client := dgo_test.GetDgraphClient()
	// Get and save buyers to the database
	storage.SaveData(client)

	// Initialize chi router
	r := chi.NewRouter()

	// Global middleware to log requests
	r.Use(middleware.Logger)

	// Route handlers for endpoints
	// Parte 1: Listar productos, compradores y transacciones
	// http://localhost:3000/home?date=1623706302&limit=200&offset=0
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {

		// type Res struct {
		// 	Buyers       []models.BuyerDgraph       `json:"buyers"`
		// 	Products     []models.ProductDgraph     `json:"products"`
		// 	Transactions []models.TransactionDgraph `json:"transactions"`
		// }

		// Create response
		reqUrl := r.URL.String()
		// var res Res
		// res.Buyers = utils.AdvancedResults(client, reqUrl)
		// res.Products = utils.AdvancedResults(client, reqUrl)
		// res.Transactions = utils.AdvancedResults(client, reqUrl)
		// Set header to send json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(utils.AdvancedResults(client, reqUrl))
	})

	// Parte 2: Listar compradores
	r.Get("/buyers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(buyers.GetBuyersFromDB(client, ""))
	})

	// Parte 3: Consultar comprador
	r.Get("/buyers/{id}", func(w http.ResponseWriter, r *http.Request) {

		// Get buyer transactions
		// Create response
		res := buyers.GetBuyer(client, chi.URLParam(r, "id"))

		// Set header to send json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	// Run server on port 3000
	fmt.Println("Listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
