package main

import (
	"sync"
)

type Resource string

type Ip int64

type Node struct {
	conn  map[Resource]Ip
	mutex sync.Mutex
}

func (n *Node) Set(r *Resource, i Ip) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.conn[*r] = i
}
func (n *Node) Get(r *Resource) Ip {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.conn[*r]
}
