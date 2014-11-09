package graph

type reachableState struct {
    graf *G
    seen []bool
    result [][]bool
}

func reachableDFS(state *reachableState, root, u int) {
    for _, v := range state.graf.Vertices(u) {
        if !state.seen[v] {
            state.seen[v] = true
            state.result[root][v] = true
            if root != v {
                reachableDFS(state, root, v)   
            }
        }
    }
}

// Given a directed graph, Reachable() returns a 2D boolean array storing
// the reachability of every pair of vertices within the graph. This
// implementation is a naive one. It simply runs a depth first search for
// every single vertice for a total complexity of O(V * (V + E)).
func Reachable(graf *G) [][]bool {
    state := &reachableState {
        graf,
        nil,
        make([][]bool, graf.V()),
    }
    for root := 0; root < graf.V(); root++ {
        state.seen = make([]bool, graf.V())
        state.result[root] = make([]bool, graf.V())
        reachableDFS(state, root, root)
    }
    return state.result
}
