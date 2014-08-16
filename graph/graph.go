/*
    The graph package works with edge-weighted directed graphs, a flexible 
    type of graphs that can be adjusted to behave like non-weighted directed
    graphs or undirected graphs. However, if you want vertice-weighted graphs,
    for example, you will have to look somwhere else.

    There are many ways one can represent edge-weighted directed graphs. For
    the sake of simplicity, we shall provide only one such representation. The
    adjacent list saves some space and performs reasonably well in graphs
    typically found in practice barring the densest ones.

    Some important graph algorithms are implemented via the following functions
    (you are recommended to read them in this order):

    - Topological sort: graph.TopoSort()
    - Cycle detection: graph.HasCycle()
    - Strongly connected components: graph.SCC()
    - Minimum spanning tree: graph.MST()
    - Shortest paths: graph.ShortestPaths()

    See: 
    1) http://algs4.cs.princeton.edu/41undirected/
    2) http://algs4.cs.princeton.edu/42directed/
    3) http://algs4.cs.princeton.edu/43mst/
    4) http://algs4.cs.princeton.edu/44sp/
    5) http://www.seas.gwu.edu/~simhaweb/alg/lectures/module7/module7.html
    6) https://www.ics.uci.edu/~eppstein/161/960220.html
    7) http://www.cs.nyu.edu/courses/summer04/G22.1170-001/6a-Graphs-More.pdf
*/

package graph

import (
    "fmt"
)

// For safety and simplicity, Edge is a read-only type.

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

// For safety and simplicity, G requires the number of vertices up-front; it
// can add edges but cannot remove them; it cannot modify the added edges.

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

func (me G) Adj(u int) []Edge {
    return me.adj[u]
}

// Add a new edge u-v. Panic if either vertice if out of range. 

func (me G) Add(u, v int, weight float32) {
    me.checkBounds(u)
    me.checkBounds(v)
    e := Edge { u, v, weight }
    me.adj[u] = append(me.adj[u], e)
    me.nE++
}

func (me G) checkBounds(u int) {
    if u < 0 || u >= me.nV {
        panic(fmt.Sprintf("vertex %d out of range [0, %d)", u, me.nV))
    }
}
