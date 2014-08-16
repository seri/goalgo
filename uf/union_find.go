package uf

type UnionFind interface {
    Reset(size int)            // construct a new non-directed graph of this size
    Size() int                 // number of vertices
    Count() int                // number of connected components
    Connected(p, q int) bool   // is there a path from p to q? (a `find` request)
    Union(p, q int)            // add the edge p-q
}
