package buyers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/models"
)

type GetBuyerRes struct {
	Transactions           []models.TransactionDgraph
	Products               []models.ProductDgraph
	ProductRecommendations []models.ProductDgraph
	Buyers                 []models.BuyerDgraph
}

type Query struct {
	NewQuery []models.TransactionDgraph `json:"new_query"`
}

var res GetBuyerRes

func GetBuyer(client *dgo.Dgraph, buyerId string) GetBuyerRes {
	q := fmt.Sprintf(`{
		new_query(func: has(device)) {
		uid
		id
		ip
		device
		buyer @filter(eq(id, %s)){
	     uid
		 id
	   }
	    products{
		 id
       }
	  }
	}`, buyerId)

	ctx := context.Background()

	resp, err := client.NewTxn().Query(ctx, q)

	if err != nil {
		log.Fatal(err)
	}

	var trnsact Query
	err = json.Unmarshal(resp.Json, &trnsact)
	if err != nil {
		log.Fatal(err)
	}

	// Only get transactions that have buyer predicate
	var filteredTransactions []models.TransactionDgraph
	for _, trnsactItem := range trnsact.NewQuery {
		if trnsactItem.Buyer != nil {
			filteredTransactions = append(filteredTransactions, trnsactItem)
		}
	}

	if len(filteredTransactions) > 0 {
		// Update transactions
		res.Transactions = filteredTransactions

		// Update products in buyer transactions
		for _, trnsactItem := range res.Transactions {
			res.Products = BuyerProducts(client, trnsactItem.Products)
		}
		// Get buyers by ip and also set recommendations
		fmt.Println("BUYER IP:", res.Transactions[0].IP)
		res.Buyers = BuyersByIp(client, res.Transactions[0].IP, buyerId)
	}

	return res
}

func BuyerProducts(client *dgo.Dgraph, product_ids []*models.ProductDgraph) []models.ProductDgraph {
	var buyerProds []models.ProductDgraph
	// Search for each prod Id
	for _, prodItem := range product_ids {
		q := fmt.Sprintf(`{
			new_query(func: type(Product)) @filter(eq(id, %s)){
			id
			name
			price
		  }
		}`, prodItem.ID)

		ctx := context.Background()

		resp, err := client.NewTxn().Query(ctx, q)

		if err != nil {
			log.Fatal(err)
		}

		type ProductsQuery struct {
			NewQuery []models.ProductDgraph `json:"new_query"`
		}

		var buyerProd ProductsQuery
		err = json.Unmarshal(resp.Json, &buyerProd)
		if err != nil {
			log.Fatal(err)
		}
		buyerProds = append(buyerProds, buyerProd.NewQuery[0])
	}
	return buyerProds
}

func BuyersByIp(client *dgo.Dgraph, buyerIp string, buyerId string) []models.BuyerDgraph {
	q := fmt.Sprintf(`{
		new_query(func: type(Transaction)) @filter(eq(ip, %s) AND NOT eq(id, %s)){
			ip
			buyer{
				id
			}
			products{
				id
			}
		}
	}`, buyerIp, buyerId)

	// Products from buyers with the same IP
	ctx := context.Background()

	resp, err := client.NewTxn().Query(ctx, q)

	if err != nil {
		log.Fatal(err)
	}

	var trnsactItems Query
	err = json.Unmarshal(resp.Json, &trnsactItems)
	if err != nil {
		log.Fatal(err)
	}

	var buyersInfo []models.BuyerDgraph
	for _, item := range trnsactItems.NewQuery {
		buyersInfo = append(buyersInfo, BuyerInfo(client, item.Buyer.ID))
	}

	// Update products in buyer transactions
	for _, trnsactItem := range trnsactItems.NewQuery {
		res.ProductRecommendations = BuyerProducts(client, trnsactItem.Products)
		// Only recommend a few products
		if len(res.ProductRecommendations) >= 5 {
			break
		}
	}

	return buyersInfo
}

func BuyerInfo(client *dgo.Dgraph, buyerId string) models.BuyerDgraph {
	q := fmt.Sprintf(`{
		new_query(func: type(Buyer)) @filter(eq(id, %s)){
			id
			name
			age
		}
	}`, buyerId)

	// Products from buyers with the same IP
	ctx := context.Background()

	resp, err := client.NewTxn().Query(ctx, q)

	if err != nil {
		log.Fatal(err)
	}

	type BuyerQuery struct {
		NewQuery []models.BuyerDgraph `json:"new_query"`
	}

	var BuyerInfo BuyerQuery
	err = json.Unmarshal(resp.Json, &BuyerInfo)
	if err != nil {
		log.Fatal(err)
	}

	return BuyerInfo.NewQuery[0]
}
