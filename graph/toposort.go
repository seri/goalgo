package graph

type topoState struct {
    graf *G
    seen []bool
    result []int
    index int
}

func topoDFS(state *topoState, u int) {
    state.seen[u] = true
    for _, v := range state.graf.Vertices(u) {
        if !state.seen[v] {
            topoDFS(state, v)
        }
    }
    state.result[state.index] = u
    state.index--
}

// This function assumes that the given graph is acyclic. If you are unsure,
// consider calling HasCycle() first to check. Now, all we are doing to find
// the topological order is to start a depth first search and save the order in
// which vertices are completed. So this will run in O(V + E).
func TopoSort(graf *G) []int {
    state := &topoState {
		graf,
		make([]bool, graf.V()),
        make([]int, graf.V()),
		graf.V() - 1,
	}
    for u := 0; u < graf.V(); u++ {
        if !state.seen[u] {
            topoDFS(state, u)
        }
    }
    return state.result
}
