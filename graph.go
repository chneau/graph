package graph

import (
	"fmt"
	"strings"
)

type Graph struct {
	Vertices map[int]*Vertex
}

func (g *Graph) AddEdge(from, to, cost int) {
	if _, exist := g.Vertices[from]; !exist {
		v := NewVertex()
		v.AddEdge(to, cost)
		g.Vertices[from] = v
		return
	}
	g.Vertices[from].AddEdge(to, cost)
}

func (g *Graph) String() string {
	str := ""
	for from, vertex := range g.Vertices {
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

func New() *Graph {
	g := &Graph{
		Vertices: map[int]*Vertex{},
	}
	return g
}
