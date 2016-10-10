package main

import (
	"sync"
)

type Node struct {
	Conns map[string]int64
	Mutex sync.Mutex
}

type  New() *Node {
	Node node=Node{Conns:make(map[string]int64)}
	return &node
}