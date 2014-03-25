/*
    The uf package solves the dynamic connectivity problem. Given a graph with
    a fixed number of vertices, you are going to add edges as they come by 
    calling `union`, and you want to quickly check whether some two vertices 
    are in the same connected component via a `find` request.

    We provide two implementations of the UnionFind API, each having a different
    performance characteristic.

    The runnable client uf_client.go benchmarks them against each other.

    See: http://algs4.cs.princeton.edu/15uf/
*/

package uf

type UnionFind interface {
    Reset(size int)            // construct a new non-directed graph of this size
    Size() int                 // number of vertices
    Count() int                // number of connected components
    Connected(p, q int) bool   // is there a path from p to q? (a `find` request)
    Union(p, q int)            // add the edge p-q
}
