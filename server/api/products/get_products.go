package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/models"
)

type Root struct {
	NewQuery []models.ProductDgraph `json:"new_query"`
}

func GetProductsFromDB(client *dgo.Dgraph, date string) []models.ProductDgraph {

	// Get all instances that have a name predicate (all)
	var q string

	if date != "" {
		q = fmt.Sprintf(`{
			new_query(func: has(price), first: 20) @filter(ge(date, %s)){
				id
				name
				price
				date
			}
		}`, date)
	} else {
		q = `{
			new_query(func: has(age)) {
				id
				name
				age
				date
			}
		}`
	}

	// Perform query
	ctx := context.Background()
	resp, err := client.NewTxn().Query(ctx, q)

	if err != nil {
		log.Fatal(err)
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r.NewQuery
}
