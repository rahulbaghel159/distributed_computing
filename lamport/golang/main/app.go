package golang

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/baghelrahul159/distrbuted_computing/golang/node"
)

func main() {
	//initialise node
	node := InitNode()

	//initialise router
	router := mux.NewRouter()

	//initialise handlers
	InitHandler(router, node)

	//start local server
	log.Fatal(http.ListenAndServe(":8080", router))
}
