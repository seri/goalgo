package graph

import (
    "container/list"
    "math"
    pq "github.com/seri/goalgo/pq/mutable"
)

// A struct type holding the result of running Dijkstra against a graph and a
// given single-source.
type DijkstraResult struct {
    source int
	cost []float64
	prev []int
}

// Returns the source vertice
func (me DijkstraResult) Source() int {
    return me.source
}

// Returns the shortest distance from source to the given target.
func (me DijkstraResult) DistTo(target int) float64 {
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
func (me DijkstraResult) PathTo(target int) []int {
    l := list.New()
    for u := target; u != -1; u = me.prev[u] {
        l.PushFront(u)
    }
    return listToIntSlice(l)
}

// Run Dijkstra against the given graph and a chosen source vertice. The
// returned struct contains information on the shortest paths from the source
// to every other vertice in the graph.
func Dijkstra(graf *G, source int) *DijkstraResult {
    cost := make([]float64, graf.V())
    prev := make([]int, graf.V())
    for u := range cost {
        cost[u] = math.Inf(1)
        prev[u] = -1
    }
    cost[source] = 0

    clone := make([]float64, graf.V())
    copy(clone, cost)
    queue := pq.New(pq.Reverse(pq.Float64Slice(clone)))

    for !queue.Empty() {
        u := queue.Pop()
        for _, edge := range graf.Edges(u) {
            v := edge.To()
            newCost := edge.Weight() + cost[u]
            if newCost < cost[v] {
                cost[v] = newCost
                prev[v] = u
                queue.Change(v, newCost)
            }
        }
    }

    return &DijkstraResult { source, cost, prev }
}
