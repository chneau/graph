package graph

import (
	"sort"
)

func DijkstraShortest(g Graph, from, to int) (int, []int) {
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
			lp := len(vertices[visiting].Path)
			if _, exist := vertices[k]; !exist { // if doenst exist, add vertex
				vertices[k] = &VertexInfo{
					Distance: v + vertices[visiting].Distance,
					Path:     append(vertices[visiting].Path[:lp:lp], k),
				}
				toVisit = append(toVisit, k)
			} else {
				newDistance := v + vertices[visiting].Distance
				if vertices[k].Distance > newDistance { // update only if better
					vertices[k].Distance = newDistance
					vertices[k].Path = append(vertices[visiting].Path[:lp:lp], k)
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
