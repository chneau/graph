package graph

import (
	"sort"
)

type Vertex struct {
	Neighbours map[int]int
	Order      []int
}

func (v *Vertex) Sort() {
	sort.Slice(v.Order, func(i, j int) bool {
		return v.Neighbours[v.Order[i]] < v.Neighbours[v.Order[j]]
	})
}

func (v *Vertex) AddEdge(to, cost int) {
	if value, exist := v.Neighbours[to]; exist {
		if value == cost { // cost doesnt change, noop
			return
		}
		v.Neighbours[to] = cost // cost change and need sort
		v.Sort()
		return
	}
	v.Neighbours[to] = cost // new edge
	v.Order = append(v.Order, to)
	v.Sort()
}

func NewVertex() *Vertex {
	v := &Vertex{
		Neighbours: map[int]int{},
	}
	return v
}
