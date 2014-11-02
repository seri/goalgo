// Package sort hosts a handful of sorting algorithms.
// 
// Perhaps more interesting than these well-known algorithms are the way that the
// Go language approaches polymorphism through a flexible interface semantics. One
// can use type-alias to adapt a known type to a new interface. One can embed an
// interface into a struct for inheritance. Read code for details.
// 
// Examples:
// 
//      * https://github.com/seri/goalgo/blob/master/examples/sort_client.go
//      tests the correctness of these algorithms and benchmarks them.
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

// For convenience, we provide an adaptation of []int against the Sortable
// interface.
type IntSlice []int

func (me IntSlice) Size() int {
    return len(me)
}
func (me IntSlice) Less(i, j int) bool {
    return me[i] < me[j]
}
func (me IntSlice) Exch(i, j int) {
    me[i], me[j] = me[j], me[i]
}

// For convenience, we provide an adaptation of []string against the Sortable
// interface.
type StringSlice []string

func (me StringSlice) Size() int {
    return len(me)
}
func (me StringSlice) Less(i, j int) bool {
    return me[i] < me[j]
}
func (me StringSlice) Exch(i, j int) {
    me[i], me[j] = me[j], me[i]
}

// This trick allows one to quickly reverse the order of a Sortable
type rSortable struct {
    Sortable
}
func (me rSortable) Less(i, j int) bool {
    return me.Sortable.Less(j, i)
}

// Tweak a Sortable collection so that when a Sorter is used, the order will be
// reversed. Which means you can sort a collection in descending order like this:
//      sort.QuickSort(sort.Reverse(collection))
func Reverse(x Sortable) Sortable {
    return &rSortable { x }
}
