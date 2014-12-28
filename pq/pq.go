// Package pq provides a priority queue implementation using a maxium heap.
// It supports both push and pop operations in logarithmic time.
// 
// Examples:
// 
//      * examples/pq_exercise.go shows how you may use this package in practice.
// 
// References:
// 
//      1. http://algs4.cs.princeton.edu/24pq/
package pq

import (
    "github.com/seri/goalgo/sort"
)

// As long as a collection has a size, can compare any two elements, can swap any
// two elements, and knows how to add and remove things in a first-in-first-out
// manner (slices and doubly linked lists, for example), it can also behave as a
// priority queue. 
type Interface interface {
    sort.Sortable
    Push(x interface{})
    Pop() interface{}
}

// O(N). Convert the given collection to a priority queue.
func Heapify(a Interface) {
    n := a.Size()
    for i := n/2 - 1; i >= 0; i-- {
        sink(a, i, n)
    }
}

// O(lgN). Push an element to the priority queue.
func Push(a Interface, x interface{}) {
    a.Push(x)
    swim(a, a.Size() - 1)
}

// O(lgN). Remove the element with the highest priority from the queue and
// retrieve it.
func Pop(a Interface) interface{} {
    if a.Size() == 0 {
        panic("Cannot pop from an empty priority queue")
    }
    n := a.Size() - 1
    a.Exch(0, n)
    sink(a, 0, n)
    return a.Pop()
}

func swim(a Interface, k int) {
    for {
        i := (k - 1)/2
        if k == 0 || !a.Less(i, k) {
            break
        }
        a.Exch(i, k)
        k = i
    }
}

func sink(a Interface, k, n int) {
    for {
        i := 1 + (2 * k)
        if i >= n {
            break
        }
        if i + 1 < n && a.Less(i, i + 1) {
            i = i + 1
        }
        if a.Less(k, i) {
            a.Exch(k, i)
            k = i
        } else {
            break
        }
    }
}
