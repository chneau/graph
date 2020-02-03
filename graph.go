package graph

import (
	"fmt"
	"strings"
)

type Graph map[int]*Vertex

func (g Graph) Best(from, to int) (int, []int) {
	return 0, nil
}

func (g Graph) AddEdge(from, to, cost int) {
	if _, exist := g[from]; !exist {
		v := NewVertex()
		v.AddEdge(to, cost)
		g[from] = v
		return
	}
	g[from].AddEdge(to, cost)
}

func (g Graph) String() string {
	str := ""
	for from, vertex := range g {
		str += "("
		vertexstr := []string{}
		for _, to := range vertex.Order {
			vertexstr = append(vertexstr, fmt.Sprintf("%d->%d:%d", from, to, vertex.Neighbours[to]))
		}
		str += strings.Join(vertexstr, ",")
		str += ")"
	}
	return str
}

func New() Graph {
	g := map[int]*Vertex{}
	return g
}
