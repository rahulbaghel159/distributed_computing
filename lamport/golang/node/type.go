package node

type Node struct {
	LamportTime int
}

type StringNode struct {
	LamportTime string `json:"time"`
}
