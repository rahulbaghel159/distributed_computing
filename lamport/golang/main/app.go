package golang

import (
	"log"
	"net/http"

	node "github.com/baghelrahul159/distrbuted_computing/golang/node"
	"github.com/gorilla/mux"
)

func main() {
	//initialise node
	n := node.InitNode()

	//initialise router
	router := mux.NewRouter()

	//initialise handlers
	node.InitHandler(router, n)

	//start local server
	log.Fatal(http.ListenAndServe(":8080", router))
}
