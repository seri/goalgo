package graph

import (
    "sort"
)

type sccState struct {
    graf *G
    result [][]int
    seen []bool
    component int
}

func sccAddComponent(state *sccState) {
    state.component++
    state.result = append(state.result, make([]int, 0))
}

func sccAddVertice(state *sccState, u int) {
    state.result[state.component] = append(state.result[state.component], u)
}

func sccDFS(state *sccState, u int) {
    state.seen[u] = true
    for _, v := range state.graf.Vertices(u) {
        if !state.seen[v] {
            sccAddVertice(state, v)
            sccDFS(state, v)
        }
    }
}

// Given a directed graph, SCC() divides all its vertices into different
// components called its strongly connected components. There exists a path
// between any two vertices within each such component (hence called a strongly
// connected component). If we reduce each such component into a single vertice
// and add edges between such vertices whenever their corresponding components
// are connected, we will obtain the kernel graph of the initial graph (the 
// proof that this graph is a DAG is left as an exercise to the reader). The
// algorithm in use here is Kosaraju-Sharir's, an ingenious algorithm. Refer to
// [5] for details and proof. Complexity is O(V + E).
func SCC(graf *G) [][]int {
    state := &sccState {
        graf,
        make([][]int, 0),
        make([]bool, graf.V()),
        -1,
    }
    for _, u := range TopoSort(graf.Reverse()) {
        if !state.seen[u] {
            sccAddComponent(state)
            sccAddVertice(state, u)
            sccDFS(state, u)
        }      
    }
    for _, component := range state.result {
        sort.Sort(sort.IntSlice(component))
    }
    return state.result
}