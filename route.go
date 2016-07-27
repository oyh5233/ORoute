package main

//	"sync"

//	"github.com/oyh5233/olog"

const (
	MaxNumber = 10000000
)

type Route struct {
	Nodes [MaxNumber]*Node
}

func (r *Route) Start() {
	//reinit all Nodes
	for i := 0; i < MaxNumber; i++ {
		n := Node{conn: make(map[Resource]Ip)}
		r.Nodes[i] = &n
	}
}

func (r *Route) Stop() {

}
