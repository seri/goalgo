Package goalgo provides implementations of some fundamental algorithms, most of
which come from the Algorithms course of the university of Princeton.

* Implementations should be clear and appropriately commented.
* Each algorithm or data structure should be accompanied by a runnable example
and if possible, a comprehensive test suite.

[Package documentation](http://godoc.org/github.com/seri/goalgo)

## List of implemented algorithms

### Miscellanous

* Knuth Shuffle

### Union Find

* Quick Find
* Weighted Quick Union

### Sorting

* Selection Sort
* Insertion Sort
* Shell Sort
* Merge Sort
* Quick Sort

### Priority Queue

* Maximum Heap
* Heap Sort

### Symbol Table

* Binary Search Tree
* Left-leaning Red Black Tree

### Graph

* Topological Sorting
* Cycle Detection
* Kosaraju-Sharir's Strongly Connected Components
* Kruskal's Minimum Spanning Tree
* Dijkstra's Shortest Paths

## How to run

In case you are new to the Go ecosystem, this is a quick start to run things:

    $ cd ~
    $ mkdir goroot
    $ export GOPATH="$HOME/goroot"
    $ go install github.com/seri/goalgo
    $ cd goroot/src/github.com/seri/goalgo/examples
    $ go run sort_client.go
