package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benjimouse/timelogutil"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (*query) Timelog() string {

	currentTime := time.Now() // The current time
	n := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
	result := timelogutil.GetTasksSince(n)
	j, _ := json.Marshal(result)
	return string(j)

}

func main() {

	s := `
                schema {
                        query: Query
                }
                type Query {
                        timelog: String!
                }
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
