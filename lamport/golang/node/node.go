package node

import (
	"log"
	"net/http"
)

func InitNode() *Node {
	return &Node{LamportTime: 0}
}

//send message inclusing current lamport time to another node in case of an internal event
func (node *Node) sendMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending message")
}

//receive message from another node and update local Lamport Time
func (node *Node) receiveMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("recieving message")
}

//return latest Lamport Time for the Node
func (node *Node) lamportTime(w http.ResponseWriter, r *http.Request) {
	log.Println("lamport Time", node.LamportTime)
}
