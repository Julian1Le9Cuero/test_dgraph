package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Buyers
// type Buyer struct {
// 	ID   string `json:"id,omitempty"`
// 	Name string `json:"name,omitempty"`
// 	Age  int    `json:"age,omitempty"`
// }

type BuyerDgraph struct {
	Uid   string    `json:"uid,omitempty"`
	ID    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Age   int       `json:"age,omitempty"`
	Date  time.Time `json:"date,omitempty"`
	DType []string  `json:"dgraph.type,omitempty"`
}

// Products
// type Product struct {
// 	ID    string `json:"id,omitempty"`
// 	Name  string `json:"name,omitempty"`
// 	Price int    `json:"price,omitempty"`
// }

type ProductDgraph struct {
	Uid   string    `json:"uid,omitempty"`
	ID    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Price int       `json:"price,omitempty"`
	Date  time.Time `json:"date,omitempty"`
	DType []string  `json:"dgraph.type,omitempty"`
}

// Transactions
// type Transaction struct {
// 	ID       string          `json:"id,omitempty"`
// 	Buyer    BuyerDgraph     `json:"buyer,omitempty"`
// 	IP       string          `json:"ip,omitempty"`
// 	Device   string          `json:"device,omitempty"`
// 	Products []ProductDgraph `json:"products,omitempty"`
// 	Date     time.Time       `json:"date,omitempty"`
// }

type TransactionDgraph struct {
	Uid      string           `json:"uid,omitempty"`
	ID       string           `json:"id,omitempty"`
	Buyer    *BuyerDgraph     `json:"buyer,omitempty"`
	IP       string           `json:"ip,omitempty"`
	Device   string           `json:"device,omitempty"`
	Products []*ProductDgraph `json:"products,omitempty"`
	Date     time.Time        `json:"date,omitempty"`
	DType    []string         `json:"dgraph.type,omitempty"`
}

func (b *BuyerDgraph) setDate() {
	b.Date = time.Now()
}

func (p *ProductDgraph) setDate() {
	p.Date = time.Now()
}

func (t *TransactionDgraph) setDate() {
	t.Date = time.Now()
}

func GetBuyers(url string) []BuyerDgraph {
	resp, err := http.Get(url)
	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatalln(err)
	}

	// Close body to avoid memory leaks
	defer resp.Body.Close()
	
	// Unmarshal json
	var buyers []BuyerDgraph

	// Convert response body from JSON to Buyer slice
	err = json.Unmarshal(body, &buyers)

	if err != nil {
		panic(err)
	}

	return buyers
}

func GetProducts(url string) []*ProductDgraph {
	resp, err := http.Get(url)

	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}
	reader := csv.NewReader(resp.Body)

	reader.Comma = '\'' // Use apostrofe-delimited instead of comma

	// reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	// Slice of products
	var products []*ProductDgraph

	for _, each := range csvData {
		var product ProductDgraph
		product.ID = each[0]
		product.Name = each[1]
		product.Price, _ = strconv.Atoi(each[2])
		products = append(products, &product)
	}

	return products
}

func GetTransactions(url string) []*TransactionDgraph {
	resp, err := http.Get(url)

	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	// Close body to avoid memory leaks
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// [\s\p{Zs}]{2,}

	var transactions []*TransactionDgraph
	s := string(body)
	i := strings.Index(s, ")")
	char := s[i+1 : i+3]  // Get the double space not being recognized as string
	char2 := s[i+1 : i+2] // Get single space

	for _, transact := range strings.Split(s, char) {
		each := regexp.MustCompile(char2).Split(transact, 5)
		// If could not split line then skip the current iteration
		if len(each) != 5 {
			continue
		}

		var Transaction TransactionDgraph
		Transaction.ID = string(each[0])
		var NewBuyer BuyerDgraph
		NewBuyer.ID = string(each[1])
		NewBuyer.Uid = "_:" + NewBuyer.ID
		Transaction.Buyer = &NewBuyer
		Transaction.IP = string(each[2])
		Transaction.Device = string(each[3])

		// Convert string array of product ids to array of Products
		product_ids := strings.Split(regexp.MustCompile(`\)|\(`).ReplaceAllString(each[4], ""), ",")
		var products []*ProductDgraph

		for _, productID := range product_ids {
			var newProduct ProductDgraph
			newProduct.ID = productID
			newProduct.Uid = "_:" + newProduct.ID
			products = append(products, &newProduct)
		}

		Transaction.Products = products
		// _ = json.Unmarshal([]byte(product_ids), &Transaction.Products)

		transactions = append(transactions, &Transaction)
	}

	return transactions
}
