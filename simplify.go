package graph

import "log"

func RemoveEdge(v *Vertex, edge int) {
	delete(v.Neighbours, edge)
	for i := range v.Order {
		if v.Order[i] == edge {
			v.Order = append(v.Order[:i], v.Order[i+1:]...)
			break
		}
	}
}
func Simplify(g Graph) {
	l := len(g)
	for {
		simplify(g)
		if l == len(g) {
			break
		}
		l = len(g)
	}
	for {
		biSimplify(g)
		if l == len(g) {
			break
		}
		l = len(g)
	}
}
func simplify(g Graph) {
	optimisable := map[int]bool{}
	where := map[int][]int{}
	for k, v := range g {
		if len(v.Order) == 1 { // if edge only going to one vertix
			optimisable[k] = true
		}
		for _, i := range v.Order { // map where a vertix appear
			where[i] = append(where[i], k)
		}
	}
	for k := range optimisable {
		if len(where[k]) != 1 { // remove vertix with multiple parrents
			delete(optimisable, k)
		}
	}
	for k := range where {
		if _, exist := optimisable[k]; !exist { // remove useless data on where map
			delete(where, k)
		}
	}
	for mid := range optimisable {
		from := where[mid][0]
		if mid == from { // a round has been reduced
			RemoveEdge(g[from], mid)
			continue
		}
		if _, exist := g[from]; !exist {
			continue
		}
		to := g[mid].Order[0]
		ok := simplifyVertices(g, from, mid, to)
		if !ok {
			log.Println("NOT OK FOR", from, mid, to)
		}
		delete(g, mid)
	}
}

//
func biSimplify(g Graph) {
	optimisable := map[int]bool{}
	where := map[int][]int{}
	for k, v := range g {
		if len(v.Order) == 2 { // if edge only going to two vertices
			optimisable[k] = true
		}
		for _, i := range v.Order { // map where a vertex appear
			where[i] = append(where[i], k)
		}
	}
	for k := range optimisable {
		if len(where[k]) != 2 { // remove vertex with multiple parrents
			delete(optimisable, k)
		}
	}
	for k := range where {
		if _, exist := optimisable[k]; !exist { // remove useless data on where map
			delete(where, k)
		}
	}
	for mid := range optimisable {
		from := where[mid][0]
		to := where[mid][1]
		if _, exist := g[from]; !exist {
			continue
		}
		if _, exist := g[to]; !exist {
			continue
		}
		if _, exist := g[mid]; !exist {
			continue
		}
		if _, exist := g[from].Neighbours[mid]; !exist {
			continue
		}
		if _, exist := g[mid].Neighbours[to]; !exist {
			continue
		}
		if _, exist := g[to].Neighbours[mid]; !exist {
			continue
		}
		if _, exist := g[mid].Neighbours[from]; !exist {
			continue
		}
		if from == mid {
			// RemoveEdge(g[from], mid)
			continue
		}
		if mid == to {
			// RemoveEdge(g[mid], to)
			continue
		}
		ok := simplifyVertices(g, from, mid, to)
		if !ok {
			log.Println("1 NOT OK FOR", from, mid, to)
		}
		ok = simplifyVertices(g, to, mid, from)
		if !ok {
			log.Println("2 NOT OK FOR", from, mid, to)
		}
		delete(g, mid)
	}
}

func simplifyVertices(g Graph, from, mid, to int) bool {
	if _, exist := g[from].Neighbours[mid]; !exist {
		return false
	}
	if _, exist := g[mid].Neighbours[to]; !exist {
		return false
	}
	cost1 := g[from].Neighbours[mid]
	cost2 := g[mid].Neighbours[to]
	cost := cost1 + cost2
	RemoveEdge(g[from], mid)
	g[from].AddEdge(to, cost)
	return true
}
