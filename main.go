// The inCommon package is a collaborative-filtering recommendation engine
// making use of the graph and NoSQL models of Neo4j
package main

import (
	"inCommon/api"
	"log"
	"net/http"
)

// Starts HTTP server and router
func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
