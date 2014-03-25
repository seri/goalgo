/*
    Package sort hosts a handful of sorting algorithms abstracted as `Sorter`s.

    Perhaps more interesting than these well-known algorithms are the way that
    the Go language approaches polymorphism through a flexible interface
    semantics. One can use type-alias to adapt a known type to a new interface.
    One can embed an interface into a struct for inheritance.

    The runnable sort_client.go tests these algorithms and benchmarks them 
    against each other.

    See:
    1) http://algs4.cs.princeton.edu/21elementary/
    2) http://algs4.cs.princeton.edu/22mergesort/
    3) http://algs4.cs.princeton.edu/23quicksort/
*/

package sort

type Sortable interface {
    Size() int
    Less(i, j int) bool
    Exch(i, j int)
}

// A Sorter is a function that sorts a Sortable
type Sorter func(Sortable)

// A Cheater is very lame as it can only sort an int slice
type Cheater func([]int)


// Some popular sortable types for convenience's sake

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

func Reverse(x Sortable) Sortable {
    return &rSortable { x }
}