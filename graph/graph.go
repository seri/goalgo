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
