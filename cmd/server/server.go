package main

import (
	"net/http"

	"github.com/vaport05/gql"
)

func main() {
	http.HandleFunc("/query", gql.GqlHandler)
}
