# timelog
Playing around with go, decided to create a tool for logging time.
---
You'll want a file called `config/config.development.json` that has the details of your mongo database.
See `config/config.example.json` as an example.

## Dependencies
https://github.com/benjimouse/timelogutil

## Running as a graphql server
`go run graphql.go`
` curl -XPOST -d '{"query": "{ timelog }"}' localhost:8080/query`
