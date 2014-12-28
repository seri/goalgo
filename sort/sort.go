// Package sort hosts a handful of sorting algorithms.
// 
// Perhaps more interesting than these well-known algorithms are the way that the
// Go language approaches polymorphism through a flexible interface semantics. One
// can use type-alias to adapt a known type to a new interface. One can embed an
// interface into a struct for inheritance. Read code for details.
// 
// Examples:
// 
//      * examples/sort_client.go tests the correctness of these algorithms and 
//      benchmarks them.
// 
// References:
// 
//      1. http://algs4.cs.princeton.edu/21elementary/
//      2. http://algs4.cs.princeton.edu/22mergesort/
//      3. http://algs4.cs.princeton.edu/23quicksort/
package sort

// To use the algorithms in this package, you must first adapt your collection to
// this interface.
type Sortable interface {
    Size() int           // number of elements in the collection
    Less(i, j int) bool  // is element at index i less than element at index j
    Exch(i, j int)       // swap the two elements at the given indices
}

// A Sorter is a function that can sort a Sortable.
type Sorter func(Sortable)

// A Cheater is very lame as it can only sort an int slice.
type Cheater func([]int)
