package graph

type cycleState struct {
    g *G
    seen, done []bool
}

func cycleDFS(state *cycleState, u int) bool {
    state.seen[u] = true
    for _, edge := range state.g.Adj(u) {
        v := edge.To()
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

func HasCycle(g *G) bool {
    seen := make([]bool, g.V())
    done := make([]bool, g.V())
    state := &cycleState { g, seen, done }
    for u := 0; u < g.V(); u++ {
        if cycleDFS(state, u) {
            return true
        }
    }
    return false
}
