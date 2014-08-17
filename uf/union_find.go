// Package uf solves the dynamic connectivity problem. Given an undirected
// graph with a fixed number of vertices, you are going to add edges as they come
// by calling union, and you want to quickly check whether some two vertices are
// in the same connected component via a find request.
// 
// We provide two implementations of union find, each having a different
// performance characteristic: quick find and quick union.
// 
// Examples:
// 
//		* https://github.com/seri/goalgo/blob/master/examples/uf_client.go
//	    benchmarks the two union find solvers against each other.
// 
// References: 
// 
//		1. http://algs4.cs.princeton.edu/15uf/
package uf

// The union find API that implementations must adapt.
type UnionFind interface {
    Reset(size int)            // construct a new non-directed graph of this size
    Size() int                 // number of vertices
    Count() int                // number of connected components
    Connected(p, q int) bool   // is there a path from p to q? (aka a find request)
    Union(p, q int)            // add the edge p-q
}
