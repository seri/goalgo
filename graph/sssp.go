package graph

import (
    "container/list"
    "math"
    pq "github.com/seri/goalgo/pq/mutable"
)


// SSSP stands for single-source shortest paths. It stores the shortest paths
// from a chosen vertice (called the source) to every other vertice of a graph.
type SSSP struct {
    cost []float64
    prev []int
}

func newSSSP(nV int) *SSSP {
    cost := make([]float64, nV)
    prev := make([]int, nV)
    for u := range cost {
        cost[u] = math.Inf(1)
        prev[u] = -1
    }
    return &SSSP { cost, prev }
}

// Returns the shortest distance from source to the given target.
func (me SSSP) DistTo(target int) float64 {
    return me.cost[target]
}

func listToIntSlice(l *list.List) []int {
    slice := make([]int, 0)
    for e := l.Front(); e != nil; e = e.Next() {
        slice = append(slice, e.Value.(int))
    }
    return slice
}

// Lists all vertices in the shortest path from source to the given target.
func (me SSSP) PathTo(target int) []int {
    l := list.New()
    for u := target; u != -1; u = me.prev[u] {
        l.PushFront(u)
    }
    return listToIntSlice(l)
}

func (me *SSSP) relax(edge Edge) {
    u, v := edge.From(), edge.To()
    newCost := edge.Weight() + me.cost[u]
    if newCost < me.cost[v] {
        me.cost[v] = newCost
        me.prev[v] = u
    }
}


// Worst time O(ElogV). Space O(V). Dijkstra produces single-source shortest
// paths for an edge-weighted directed graph, but only for nonnegative weights.
// The algorithm itself is quite simple. The bulk of the work is done by a
// mutable priority queue.
//
// The algorithm works because we always visit the vertices in the order of how
// close they are to the source.
func Dijkstra(graf *G, source int) *SSSP {
    sssp := newSSSP(graf.V())
    sssp.cost[source] = 0

    clone := make([]float64, graf.V())
    copy(clone, sssp.cost)
    queue := pq.New(pq.Reverse(pq.Float64Slice(clone)))

    for !queue.Empty() {
        u := queue.Pop()
        for _, edge := range graf.Edges(u) {
            v := edge.To()
            newCost := edge.Weight() + sssp.cost[u]
            if newCost < sssp.cost[v] {
                sssp.cost[v] = newCost
                sssp.prev[v] = u
                queue.Change(v, newCost)
            }
        }
    }

    return sssp
}


// O(V + E). AcyclicSP() simply visits vertices in topological order. Recall
// that TopoSort() has complexity O(V + E), which makes AcyclicSP() linear time
// as well. Note that AcylicSP() can handle negative weights.
//
// The algorithm works because by the time we visit vertice u, we are guaranteed
// by topological order that there are no edges from u to any vertices already
// visited.
func AcyclicSP(graf *G, source int) *SSSP {
    sssp := newSSSP(graf.V())
    sssp.cost[source] = 0
    for _, u := range TopoSort(graf) {
        for _, edge := range graf.Edges(u) {
            sssp.relax(edge)
        }
    }
    return sssp
}


// Time O(V * E), space O(V) in the worst case. BellmanFord() is the most
// general solution to the single-source shortest paths problem. It can handle
// every graph without negative cycles. (The algorithm can be extended to detect
// negative cycles but we will not get into that here.)
//
// BellmandFord() makes V passes, visiting every edge in any order in each pass.
// The algorithm works because in the k-th pass, we have already found all
// shortest paths that contain at most k edges. This in turn can be proved by
// induction.
func BellmanFord(graf *G, source int) *SSSP {
    sssp := newSSSP(graf.V())
    sssp.cost[source] = 0
    for i := 0; i < graf.V(); i++ {
        for u := 0; u < graf.V(); u++ {
            for _, edge := range graf.Edges(u) {
                sssp.relax(edge)
            }
        }
    }
    return sssp
}
