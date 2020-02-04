package graph

import (
	"fmt"
	"sort"
	"strings"
)

type Graph map[int]*Vertex

func (g Graph) Best(from, to int) (int, []int) {
	type VertexInfo struct {
		Distance int
		Path     []int
	}
	vertices := map[int]*VertexInfo{
		from: &VertexInfo{Path: []int{from}},
	}
	visited := map[int]bool{}
	toVisit := []int{from}
	found := false
	for len(toVisit) > 0 && !found {
		visiting := toVisit[0] // take the first to visit
		for _, k := range g[visiting].Order {
			if visited[k] { // don't visit fully visited
				continue
			}
			v := g[visiting].Neighbours[k]
			if _, exist := vertices[k]; !exist { // if doenst exist, add vertex
				vertices[k] = &VertexInfo{Distance: v + vertices[visiting].Distance, Path: append(vertices[visiting].Path, k)}
				toVisit = append(toVisit, k)
			} else {
				newDistance := v + vertices[visiting].Distance
				if vertices[k].Distance > newDistance { // update only if better
					vertices[k].Distance = newDistance
					vertices[k].Path = append(vertices[visiting].Path[:len(vertices[visiting].Path):len(vertices[visiting].Path)], k)
				}
			}
		}
		if visiting == to {
			break
		}
		toVisit = toVisit[1:]                     // remove from tovisit
		visited[visiting] = true                  // mark visited
		sort.Slice(toVisit, func(i, j int) bool { // sort tovisit
			return vertices[toVisit[i]].Distance < vertices[toVisit[j]].Distance
		})
	}
	if _, exist := vertices[to]; !exist {
		return 0, nil
	}
	return vertices[to].Distance, vertices[to].Path
}

func (g Graph) AddBiEdge(from, to, cost int) {
	g.AddEdge(from, to, cost)
	g.AddEdge(to, from, cost)
}

func (g Graph) AddEdge(from, to, cost int) {
	if _, exist := g[to]; !exist {
		g[to] = NewVertex()
	}
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
