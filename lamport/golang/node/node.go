package node

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bitly/go-simplejson"
)

func InitNode() *Node {
	return &Node{LamportTime: 0}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//send message inclusing current lamport time to another node in case of an internal event
func (node *Node) sendMessage(w http.ResponseWriter, r *http.Request) {
	//fetching host parameter
	host, ok := r.URL.Query()["host"]

	if !ok || len(host[0]) < 1 {
		log.Println("Url Param 'host' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	str_host := host[0]

	//fetching port parameter
	port, ok := r.URL.Query()["port"]

	if !ok || len(port[0]) < 1 {
		log.Println("Url Param 'port' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	str_port := port[0]

	log.Println(node.LamportTime)
	log.Println("http://" + str_host + ":" + str_port + "/receive?" + "time=" + strconv.Itoa(node.LamportTime))

	resp, err := http.Get("http://" + str_host + ":" + str_port + "/receive?" + "time=" + strconv.Itoa(node.LamportTime))
	if resp == nil || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("status code", resp.StatusCode)
		w.WriteHeader(http.StatusBadRequest)
	}

	recievedNode := new(StringNode)
	err = json.NewDecoder(resp.Body).Decode(recievedNode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	//returning lamport time of the node
	json := simplejson.New()
	json.Set("time", recievedNode.LamportTime)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

//receive message from another node and update local Lamport Time
func (node *Node) receiveMessage(w http.ResponseWriter, r *http.Request) {
	//fetching time parameter
	time, ok := r.URL.Query()["time"]

	if !ok || len(time[0]) < 1 {
		log.Println("Url Param 'time' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	str_received_lamport_time := time[0]

	received_lamport_time, err := strconv.Atoi(str_received_lamport_time)
	if err != nil {
		log.Println("Url Param 'time' is incorrect")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//comparing time with local lamport time
	node.LamportTime = max(received_lamport_time, node.LamportTime)

	//returning lamport time of the node
	json := simplejson.New()
	json.Set("time", strconv.Itoa(node.LamportTime))

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

//return latest Lamport Time for the Node
func (node *Node) lamportTime(w http.ResponseWriter, r *http.Request) {
	//returning lamport time of the node
	json := simplejson.New()
	json.Set("time", strconv.Itoa(node.LamportTime))

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
