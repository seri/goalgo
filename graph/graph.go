// Package graph implements fundamental graph algorithms. There are many types
// of graphs but we will only deal with edge-weighted directed graphs
// represented by adjacent lists, which are arguably the most practical of all.
//
// Here are the algorithms included in the package:
//
//      1. Sort the vertices in topological order
//      2. Test if the graph is acyclic
//      3. Divide the graph into strongly connected components
//      4. Construct the minimum spanning tree
//      5. Find shortest paths between any two vertices
//
// References:
//
//      1. http://algs4.cs.princeton.edu/40graphs/
//      2. http://algs4.cs.princeton.edu/42directed/
//      3. http://algs4.cs.princeton.edu/43mst/
//      4. http://algs4.cs.princeton.edu/44sp/
//      5. http://www.seas.gwu.edu/~simhaweb/alg/lectures/module7/module7.html
//
package graph

import (
    "fmt"
)

// For simplicity, Edge is a read-only type. You can't even instantiate it.
type Edge struct {
    from, to int
    weight float32
}

func (me Edge) From() int {
    return me.from
}

func (me Edge) To() int {
    return me.to
}

func (me Edge) Weight() float32 {
    return me.weight
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
func (me G) Add(u, v int, weight float32) {
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
