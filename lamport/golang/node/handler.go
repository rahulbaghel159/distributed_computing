package node

import (
	"github.com/gorilla/mux"
)

func InitHandler(router *mux.Router, node *Node) {
	router.HandleFunc("/send", node.sendMessage)
	router.HandleFunc("/receive", node.receiveMessage)
	router.HandleFunc("/lamport", node.lamportTime)
}
