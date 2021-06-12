package buyers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/julian1le9cuero/dgraph_challenge/database/server/models"
)

type Root struct {
	NewQuery []models.BuyerDgraph `json:"new_query"`
}

func GetBuyersFromDB(client *dgo.Dgraph, date string) []models.BuyerDgraph {

	// Get all instances that have a name predicate (all)
	var q string

	if date != "" {
		q = fmt.Sprintf(`{
			new_query(func: has(age), first: 20) @filter(ge(date, %s)){
				id
				name
				age
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
