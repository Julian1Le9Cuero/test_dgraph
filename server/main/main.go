package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julian1le9cuero/dgraph_challenge/server/api/buyers"
	"github.com/julian1le9cuero/dgraph_challenge/server/api/products"
	"github.com/julian1le9cuero/dgraph_challenge/server/api/transactions"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/dgo_test"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/models"
	// "github.com/julian1le9cuero/dgraph_challenge/server/database/storage"
)

func main() {
	client := dgo_test.GetDgraphClient()
	// // Get and save buyers to the database
	// storage.SaveData(client)

	// Initialize chi router
	r := chi.NewRouter()

	// Global middleware to log requests
	r.Use(middleware.Logger)

	// Route handlers for endpoints
	// Parte 1: Listar productos, compradores y transacciones

	r.Get("/{date}", func(w http.ResponseWriter, r *http.Request) {
		// Get date from endpoint
		date_param := chi.URLParam(r, "date")

		// Date to filter
		var dateFilter string

		// Check if param has at least a number
		matched, _ := regexp.MatchString(`[0-9]+`, date_param)

		if matched {
			// Filtrar por fecha

			// Convert date string to int64
			unix_date, _ := strconv.ParseInt(date_param, 10, 64)
			// Unix Timestamp to time.Time
			timeT := time.Unix(unix_date, 0)

			dateFilter = timeT.Format("2006-01-02")

			// http.Error(w, http.StatusText(400), 400)
		} else {
			// Devolver la fecha actual por defecto
			dateFilter = time.Now().Format("2006-01-02")
		}

		type Res struct {
			Buyers       []models.BuyerDgraph       `json:"buyers"`
			Products     []models.ProductDgraph     `json:"products"`
			Transactions []models.TransactionDgraph `json:"transactions"`
		}

		// Create response
		var res Res
		res.Buyers = buyers.GetBuyersFromDB(client, dateFilter)
		res.Products = products.GetProductsFromDB(client, dateFilter)
		res.Transactions = transactions.GetTransactionsFromDB(client, dateFilter)
		// Set header to send json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
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
