package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/dgraph-io/dgo/v200"
	"github.com/julian1le9cuero/dgraph_challenge/server/database/models"
)

type Root struct {
	BuyerSlice   []models.BuyerDgraph       `json:"buyer_query"`
	ProductSlice []models.ProductDgraph     `json:"product_query"`
	TrnsactSlice []models.TransactionDgraph `json:"trnsact_query"`
}

func AdvancedResults(client *dgo.Dgraph, reqUrl string) *Root {
	u, err := url.Parse(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	urlQuery := u.Query()

	// date, offset and limit
	// Filtrar por fecha
	// Convert date string to int64
	unix_date, _ := strconv.ParseInt(urlQuery["date"][0], 10, 64)
	// Unix Timestamp to time.Time
	timeT := time.Unix(unix_date, 0)
	timeT2 := time.Unix(unix_date+24*60*60, 0)
	dateFilter := timeT.Format("2006-01-02")
	dateFilterTomorrow := timeT2.Format("2006-01-02")

	// Get all instances that have a specific type predicate (all)
	var qName string
	var q string

	model := urlQuery["model"][0]
	switch model {
	case "Product":
		qName = "product_query"
		q = fmt.Sprintf(`{
			%s(func: type(%s), first: %s, offset: %s, orderasc: id) @filter(ge(date, %s) AND lt(date, %s)){
				id
				name
				price
				date
			}
		}`, qName, model, urlQuery["limit"][0], urlQuery["offset"][0], dateFilter, dateFilterTomorrow)
	case "Transaction":
		qName = "trnsact_query"
		q = fmt.Sprintf(`{
			%s(func: type(%s), first: %s, offset: %s, orderasc: id) @filter(ge(date, %s) AND lt(date, %s)){
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
		}`, qName, model, urlQuery["limit"][0], urlQuery["offset"][0], dateFilter, dateFilterTomorrow)
	default:
		qName = "buyer_query"
		q = fmt.Sprintf(`{
			%s(func: type(%s), first: %s, offset: %s, orderasc: id) @filter(ge(date, %s) AND lt(date, %s)){
				id
				name
				age
				date
			}
		}`, qName, model, urlQuery["limit"][0], urlQuery["offset"][0], dateFilter, dateFilterTomorrow)
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

	return &r
}
