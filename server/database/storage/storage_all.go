package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/models"
)

const buyers_url string = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
const products_url string = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
const transactions_url string = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"

func SaveData(client *dgo.Dgraph) {
	// Initialize dgraph client
	// client := dgo_test.GetDgraphClient()

	// Drop/delete all the data, and start from a clean slate
	op := &api.Operation{DropAll: true}
	op.Schema = `
		id: string @index(exact) .
		name: string @index(exact) .
		age: int .
		price: int .
		buyer: uid  .
		ip: string .
		device: string .
		products: [uid]  .
		date: datetime .

		type Buyer {
			id:  string 
			name: string    
			age:  int       
			date: datetime
		}
		
		type Product {
			id:    string   
			name:  string   
			price: int      
			date: datetime 
		}
		
		type Transaction {
			id:       string          
			buyer:    Buyer @hasInverse(field: id)
			ip:       string          
			device:   string          
			products: [Product] @hasInverse(field: id)
			date: datetime @search      
		}
	`

	ctx := context.Background()
	if err := client.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

	// Get buyers from API
	buyers := models.GetBuyers(buyers_url)

	// Loop trough buyers and save each one to db
	fmt.Println("Inserting buyers...")
	for _, buyer := range buyers[0:600] {
		// Add current date to buyer
		models.FillDefault(&buyer)
		buyer.Uid = "_:" + buyer.ID
		buyer.DType = []string{"Buyer"}

		// Convert from Buyer instance to json
		pb, err := json.Marshal(buyer)

		if err != nil {
			log.Fatal(err)
		}
		mu := &api.Mutation{CommitNow: true}
		// Use JSON and pass the node that will be added
		mu.SetJson = pb

		// Run transaction through mutation (add/remove data) on db
		_, err2 := client.NewTxn().Mutate(ctx, mu)

		if err2 != nil {
			log.Fatal(err2)
		}

	}

	// Get products from API
	products := models.GetProducts(products_url)
	// Loop trough products and save each one to db
	fmt.Println("Inserting products...")
	for _, product := range products {

		// Add current date to product
		models.FillDefault(product)

		// Add Uid
		product.Uid = "_:" + product.ID
		product.DType = []string{"Product"}

		// Convert from Product instance to json
		pb, err := json.Marshal(product)

		if err != nil {
			log.Fatal(err)
		}

		mu := &api.Mutation{CommitNow: true}
		// Use JSON and pass the node that will be added
		mu.SetJson = pb

		// Run transaction through mutation (add/remove data) on db
		_, err2 := client.NewTxn().Mutate(ctx, mu)

		if err2 != nil {
			log.Fatal(err2)
		}

	}

	// Get transactions from API
	transactions := models.GetTransactions(transactions_url)
	// Loop trough transactions and save each one to db
	fmt.Println("Inserting transactions...")
	for _, transaction := range transactions[0:1600] {

		// Add current date to transaction
		models.FillDefault(transaction)
		transaction.Uid = "_:" + transaction.ID
		transaction.DType = []string{"Transaction"}

		// Convert from Transaction instance to json
		pb, err := json.Marshal(transaction)

		if err != nil {
			log.Fatal(err)
		}
		mu := &api.Mutation{CommitNow: true}
		// Use JSON and pass the node that will be added
		mu.SetJson = pb

		// Run transaction through mutation (add/remove data) on db
		_, err2 := client.NewTxn().Mutate(ctx, mu)

		if err2 != nil {
			log.Fatal(err2)
		}

	}

	fmt.Println("DONE.")
}
