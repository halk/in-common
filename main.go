// The in-common package is a data-agnostic, graph-based recommendation engine
//
// see https://github.com/halk/in-common for details
package main

import (
	"github.com/halk/in-common/api"
	"log"
	"net/http"
)

// Starts HTTP server and router
func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
