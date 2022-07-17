package main

import (
	"net/http"

	"github.com/vapor05/gql/pkg/graphql"
)

func main() {
	http.HandleFunc("/query", graphql.GqlHandler)
}
