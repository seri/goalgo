package graph

import (
    "sort"
)

type sccState struct {
    graf *G
    seen []bool
    result [][]int
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

// Given a directed graph, StrongComponents() divides all its vertices into 
// different components called its strongly connected components. There exists 
// a path between any two vertices within each such component (hence called a
// strongly connected component). If we reduce each such component into a
// single vertice and add edges between such vertices whenever their
// corresponding components are connected, we will obtain the kernel graph of
// the initial graph (the proof that this graph is a DAG is left as an exercise
// to the reader). 
// 
// The algorithm in use here is Kosaraju-Sharir's, an ingenious algorithm. We
// first run a depth first search on the reverse graph of the original graph
// and record the completion order. (This is, of course, the same as running a
// topological sort, only the result is technically not a topological order
// becase the graph may have cycles.) We then run another depth first search on
// the original graph in the order that was recorded. For every edge u->v that
// is explored in the second depth first search, we prove that there must be
// a path from v to u as well, which makes u and v strongly connected. Refer to
// [5] for proof. Complexity is O(V + E).
func StrongComponents(graf *G) [][]int {
    state := &sccState {
        graf,
        make([]bool, graf.V()),
        make([][]int, 0),
        -1,
    }
    for _, u := range TopoSort(Reverse(graf)) {
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

// This is a convenient function to convert the result of StrongComponents()
// into 2D boolean array storing whether or not any pair of vertices are in a
// strongly connected component. You have to run StrongComponents() first.
// Requires V * V space.
func InSameComponent(graf *G, scc [][]int) [][]bool {
    result := make([][]bool, graf.V())
    for u := 0; u < graf.V(); u++ {
        result[u] = make([]bool, graf.V())
    }
    for _, a := range scc {
        if len(a) == 1 {
            continue
        }
        for _, u := range a {
            for _, v := range a {
                result[u][v] = true
                result[v][u] = true
            }
        }
    }
    return result
}
