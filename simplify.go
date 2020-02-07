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
		simplifyVertices(g, from, mid, to)
		delete(g, mid)
	}
}
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
		ok := simplifyVertices(g, from, mid, to)
		if !ok {
			log.Println(ok)
		}
		ok = simplifyVertices(g, to, mid, from)
		if !ok {
			log.Println(ok)
		}
		delete(g, mid)
	}
}

func simplifyVertices(g Graph, f, m, t int) bool {
	if _, exist := g[f].Neighbours[m]; !exist {
		return false
	}
	if _, exist := g[m].Neighbours[t]; !exist {
		return false
	}
	cost1 := g[f].Neighbours[m]
	cost2 := g[m].Neighbours[t]
	cost := cost1 + cost2
	RemoveEdge(g[f], m)
	g[f].AddEdge(t, cost)
	return true
}
