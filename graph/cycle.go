package graph

type cycleState struct {
    graf *G
    seen, done []bool
}

func cycleDFS(state *cycleState, u int) bool {
    state.seen[u] = true
    for _, v := range state.graf.Vertices(u) {
        if !state.seen[v] {
            if cycleDFS(state, v) {
                return true
            }
        } else {
            if !state.done[v] {
                return true
            }
        }
    }
    state.done[u] = true
    return false
}

// Check whether the given graph has a cycle. If the function returns false, we
// call the given graph a DAG (directed acyclic graph). As this implementation
// runs one pass of depth first search, the complexity is O(V + E).
func HasCycle(graf *G) bool {
    state := &cycleState {
        graf,
        make([]bool, graf.V()),
        make([]bool, graf.V()),
    }
    for u := 0; u < graf.V(); u++ {
        if !state.seen[u] && cycleDFS(state, u) {
            return true
        }
    }
    return false
}
