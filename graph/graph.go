// Package graph implements fundamental graph algorithms. There are many types
// of graphs but we will only deal with edge-weighted directed graphs
// represented by adjacent lists.
// 
// If you do not want the edges to have weights, simply assign a unit weight of
// one to every edge. If you want an undirected graph, simply add two edges in
// both directions for every pair of vertices.
//
// Here are the algorithms included in the package:
//
//      1. Determine whether one vertice is reachable from another
//      2. Sort the vertices in topological order
//      3. Test if the graph is acyclic
//      4. Divide the graph into strongly connected components
//      5. Construct the minimum spanning tree
//      6. Find single-source shortest paths
//
// Examples
//      
//      * examples/toposort_example.go
//      * examples/cycle_example.go
//      * examples/scc_example.go
//      * examples/kruskal_example.go
//      * examples/dijkstra_example.go
//      * examples/sssp_client.go
//
// References:
//
//      1. http://algs4.cs.princeton.edu/40graphs/
//      2. http://algs4.cs.princeton.edu/42directed/
//      3. http://algs4.cs.princeton.edu/43mst/
//      4. http://algs4.cs.princeton.edu/44sp/
//      5. http://www.seas.gwu.edu/~simhaweb/alg/lectures/module7/module7.html
//      6. http://www.seas.gwu.edu/~simhaweb/alg/lectures/module8/module8.html
//      7. http://people.qc.cuny.edu/faculty/christopher.hanusa/courses/634sp12/Documents/KruskalProof.pdf
package graph

import (
    "fmt"
)

// For simplicity, Edge is a read-only type. You can't even instantiate it.
type Edge struct {
    from, to int
    weight float64
}

func (me Edge) From() int {
    return me.from
}
func (me Edge) To() int {
    return me.to
}
func (me Edge) Weight() float64 {
    return me.weight
}
func (me Edge) String() string {
    return fmt.Sprintf("%d->%d (%f)", me.from, me.to, me.weight)
}

type EdgeSlice []Edge
func (me EdgeSlice) Size() int {
    return len(me)
}
func (me EdgeSlice) Less(i, j int) bool {
    return me[i].Weight() < me[j].Weight()
}
func (me EdgeSlice) Exch(i, j int) {
    me[i], me[j] = me[j], me[i]
}

// For simplicity, G requires the number of vertices up-front; it can add edges
// but cannot remove them; it cannot modify the edges that have been added.
type G struct {
    nV int
    nE int
    adj [][]Edge
}

// Create a new graph with a fixed number of vertices
func New(nV int) *G {
    return &G { nV, 0, make([][]Edge, nV) }
}

// Return the number of vertices
func (me G) V() int {
    return me.nV
}

// Return the total number of edges
func (me G) E() int {
    return me.nE
}

// Return a slice of edges starting with vertice u
func (me G) Edges(u int) []Edge {
    return me.adj[u]
}

// Return a slice of vertices which have edges connecting to them from u
func (me G) Vertices(u int) []int {
    edges := me.adj[u]
    vertices := make([]int, len(edges))
    for i, edge := range edges {
        vertices[i] = edge.To()
    }
    return vertices
}

// Add a new edge u-v. Panic if either vertice is out of range. 
func (me G) Add(u, v int, weight float64) {
    me.checkBounds(u)
    me.checkBounds(v)
    edge := Edge { u, v, weight }
    me.adj[u] = append(me.adj[u], edge)
    me.nE++
}

func (me G) checkBounds(u int) {
    if u < 0 || u >= me.nV {
        panic(fmt.Sprintf("vertex %d out of range [0, %d)", u, me.nV))
    }
}

// Obtain a copy of the graph with all edges reversed in direction.
func Reverse(g *G) *G {
    reversed := New(g.V())
    for u := 0; u < g.V(); u++ {
        for _, edge := range g.Edges(u) {
            reversed.Add(edge.To(), edge.From(), edge.Weight())
        }
    }
    return reversed
}

// Obtain a list of every edge in an undirected graph.
func AllEdges(g *G) []Edge {
    edges := make([]Edge, g.E() / 2)
    for u := 0; u < g.V(); u++ {
        for _, edge := range g.Edges(u) {
            v := edge.To()
            if v >= u {
                edges = append(edges, edge)
            }
        }
    }
    return edges
}
