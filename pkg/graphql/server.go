package graphql

import (
	"fmt"
	"net/http"
)

func GqlHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "graphql respnse!\n")
}
