package graph

import (
	"fmt"
	"strings"
)

type Graph map[int]*Vertex

func (g *Graph) AddBiEdge(from, to, cost int) {
	g.AddEdge(from, to, cost)
	g.AddEdge(to, from, cost)
}

func (g Graph) AddEdge(from, to, cost int) {
	if _, exist := g[to]; !exist {
		g[to] = NewVertex()
	}
	if _, exist := g[from]; !exist {
		g[from] = NewVertex()
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
