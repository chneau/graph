package graph

import (
	"container/list"
)

type VertexInfo struct {
	ID       int
	Distance int
	Path     []int
}

type DLList struct {
	*list.List
}

func (l *DLList) InsertOrdered(v *VertexInfo) {
	if l.Len() == 0 {
		l.PushFront(v)
		return
	}
	back := l.Back()
	if back.Value.(*VertexInfo).Distance < v.Distance {
		l.InsertAfter(v, back)
		return
	}
	current := l.Front()
	for current.Value.(*VertexInfo).Distance < v.Distance && current.Value.(*VertexInfo).ID != v.ID {
		next := current.Next()
		if next == nil {
			break
		}
		current = next
	}
	if current.Value.(*VertexInfo).ID == v.ID {
		return
	}
	l.InsertAfter(v, current)
}

func (l *DLList) PopFront() *VertexInfo {
	e := l.Front()
	l.Remove(e)
	return e.Value.(*VertexInfo)
}

func DijkstraShortest(g Graph, from, to int) (int, []int) {
	vertices := map[int]*VertexInfo{
		from: &VertexInfo{ID: from, Path: []int{from}},
	}
	visited := map[int]bool{}
	toVisit := DLList{List: list.New()}
	toVisit.InsertOrdered(vertices[from])
	found := false
	for toVisit.Len() > 0 && !found {
		visiting := toVisit.PopFront() // take the first to visit
		for _, k := range g[visiting.ID].Order {
			if visited[k] { // don't visit fully visited
				continue
			}
			v := g[visiting.ID].Neighbours[k]
			lp := len(vertices[visiting.ID].Path)
			if _, exist := vertices[k]; !exist { // if doenst exist, add vertex
				vertices[k] = &VertexInfo{
					ID:       k,
					Distance: v + vertices[visiting.ID].Distance,
					Path:     append(vertices[visiting.ID].Path[:lp:lp], k),
				}
				toVisit.InsertOrdered(vertices[k])
			} else {
				newDistance := v + vertices[visiting.ID].Distance
				if vertices[k].Distance > newDistance { // update only if better
					vertices[k].Distance = newDistance
					vertices[k].Path = append(vertices[visiting.ID].Path[:lp:lp], k)
				}
			}
		}
		if visiting.ID == to {
			break
		}
		visited[visiting.ID] = true // mark visited
	}
	if _, exist := vertices[to]; !exist {
		return 0, nil
	}
	return vertices[to].Distance, vertices[to].Path
}
