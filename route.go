package main

//	"sync"

//	"github.com/oyh5233/olog"

const (
	MaxNodesCount = 1000
)

type Route struct {
	Nodes [MaxNodesCount]*Node
}

func (r *Route) New() {
	for i := 0; i < MaxNumber; i++ {
		n := Node.New()
		r.Nodes[i] = &n
	}
}
