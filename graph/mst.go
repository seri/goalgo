package graph

import (
    "github.com/seri/goalgo/uf"
    "github.com/seri/goalgo/sort"
)

// Construct a minimum spanning tree out of a weighted, connected, undirected
// graph. (Recall that we treat an undirected graph the same as a directed graph
// where each edge has a twin that goes in its reverse direction.) A spanning
// tree is a connected subtree of the original graph containing all the
// vertices and a subset of all the edges. (A subtree is an a cyclic subgraph.)
// The total cost of a spanning tree is the sum of the weights of its edges. A 
// mimimum spanning tree is one whose total cost is the smallest among all such
// trees.
//
// We will implement Kruskal's algorithm. We consider all the edges of the
// original graph in from the one with the smallest weight to the one with the
// largest weight. For each edge, we will add it to the resulting graph if it
// does not create a cycle. In the end, the resulting graph will be a valid
// minium spanning tree.
//
// How to quickly check whether adding an edge will create a cycle? Surely we
// cannot simply run HasCycle() every time an edge is considered. The solution
// is to use union find! Read [7] for proof. Complexity is O(E * log(V)).
func Kruskal(graf *G) *G {
    result := New(graf.V())
    cycleDetector := uf.NewQuickUnion(graf.V())
    edges := AllEdges(graf)
    sort.QuickSort(EdgeSlice(edges))
    for _, edge := range edges {
        u, v := edge.From(), edge.To()
        if !cycleDetector.Connected(u, v) {
            cycleDetector.Union(u, v)
            result.Add(u, v, edge.Weight())
            result.Add(v, u, edge.Weight())
        }
    }
    return result
}
