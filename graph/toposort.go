package graph

type topoState struct {
    g *G
    a []int
    b []bool
    k int
}

func topoDFS(s *topoState, u int) {
    s.b[u] = true
    for _, e := range s.g.Adj(u) {
        v := e.To()
        if !s.b[v] {
            topoDFS(s, v)
        }
    }
    s.a[s.k] = u
    s.k--
}

// This function assumes that the given graph is acyclic. Call 
// graph.IsAcyclic() beforehand if you are unsure.

func TopoSort(g *G) []int {
    a := make([]int, g.V())
    b := make([]bool, g.V())
    k := g.V() - 1
    s := &topoState { g, a, b, k }
    for u := 0; u < g.V(); u++ {
        if !b[u] {
            topoDFS(s, u)
        }
    }
    return a
}
