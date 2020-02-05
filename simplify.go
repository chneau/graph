package graph

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
		if _, exist := g[from]; !exist || mid == from { // a round has been reduced
			continue
		}
		to := g[mid].Order[0]
		cost1 := g[mid].Neighbours[to]
		cost2 := g[from].Neighbours[mid]
		cost := cost1 + cost2
		delete(g, mid)
		RemoveEdge(g[from], mid)
		g[from].AddEdge(to, cost)
	}
}
