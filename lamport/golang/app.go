package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	node "github.com/rahulbaghel159/distributed_computing/lamport/golang/node"
)

func main() {
	port := flag.String("port", "8080", "port on which program is running")

	flag.Parse()

	//initialise node
	n := node.InitNode()
	//initialise router
	router := mux.NewRouter()
	//initialise handlers
	node.InitHandler(router, n)

	//start local server
	log.Println("Starting Server on port", *port)
	log.Fatal(http.ListenAndServe(string(":"+*port), router))
}
