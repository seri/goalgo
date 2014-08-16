The uf package solves the dynamic connectivity problem. Given an undirected
graph with a fixed number of vertices, you are going to add edges as they come
by calling `union`, and you want to quickly check whether some two vertices are
in the same connected component via a `find` request.

We provide two implementations of union-find, each having a different
performance characteristic.

* QuickFind is an implementation of the UnionFind API that supports `find` in
constant time but it suffers linear complexity with regard to `union`. As it
tries doing more work as early as possible, this is so called the eager approach
to the dynamic connectivity problem.

* QuickUnion is an implementation of the UnionFind API that supports both `find`
and `union` in logarithmic time. As it delays doing work until the time comes,
it's a lazy approach to the dynamic connectivity problem.

## Example

The runnable client uf_client.go benchmarks them against each other.

## References: 

* http://algs4.cs.princeton.edu/15uf/
