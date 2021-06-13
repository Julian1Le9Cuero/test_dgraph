package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/models"
)

type Root struct {
	NewQuery []models.TransactionDgraph `json:"new_query"`
}

func GetTransactionsFromDB(client *dgo.Dgraph, date string) []models.TransactionDgraph {

	// Get all instances that have a name predicate (all)
	var q string

	if date != "" {
		q = fmt.Sprintf(`{
			new_query(func: has(device), first: 20) @filter(ge(date, %s)){
				id
				device
				date
				buyer{
					id
				}
				products{
					id
				}
			}
		}`, date)
	} else {
		q = `{
			new_query(func: has(device)) {
				id
				device
				date
				buyer{
					id
				}
				products{
					id
				}
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
