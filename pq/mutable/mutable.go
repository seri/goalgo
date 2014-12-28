// Package pq/mutable provides a fixed size mutable priority queue. You cannot
// add elements, but you can remove elements with the highest priorities one at
// a time, and you can change the priorities of elements. Each element is
// associated with a unique index, which is just its index in the initial
// container that you give to New(). This can be a bit confusing, so feel free
// to refer to graph/dijkstra.go for an example use.
package mutable

import (
    "github.com/seri/goalgo/sort"
)

// As long as a collection has a size, can compare any two elements, can swap any
// two elements, and knows how to change an element in a particular index, it can
// also behave as a mutable priority queue. A slice is obviously the ideal
// candidate.
type Interface interface {
    sort.Sortable
    Change(index int, value interface{})
}

type PriorityQueue struct {
    a Interface
    capacity int
    size int
    idx []int // idx[i] = the original index of element a[i]
    pos []int // pos[i] = the position of element with index i in a
    // We have to maintain that pos[idx[i]] = i
}

// O(N). Create a new mutable priority queue using the given argument as the
// internal container. After this operation, the container will be shuffled
// around to maintain heap order, but the initial index of each element is
// remembered so you can refer to it later.
func New(a Interface) *PriorityQueue {
    n := a.Size()
    idx := make([]int, n)
    pos := make([]int, n)
    for i := range idx {
        idx[i] = i
        pos[i] = i 
    }
    me := &PriorityQueue { a, n, n, idx, pos }
    for i := n/2 - 1; i >= 0; i-- {
        me.sink(i, n)
    }
    return me
}

// O(1). Retrieve the capacity (initial size) of the mutable priority queue.
func (me PriorityQueue) Capacity() int {
    return me.capacity
}

// O(1). Retrieve the size of the mutable priority queue.
func (me PriorityQueue) Size() int {
    return me.size
}

// O(1). Is the mutable priority queue empty?
func (me PriorityQueue) Empty() bool {
    return me.size == 0
}

// O(1). Does the mutable priority queue contain the given index?
func (me PriorityQueue) Contains(index int) bool {
    return me.pos[index] < me.size
}

// O(lgN). Remove the element with the highest priority from the queue and
// return its index.
func (me *PriorityQueue) Pop() int {
    if me.size == 0 {
        panic("Cannot pop from an empty priority queue")
    }
    n := me.size - 1
    ret := me.idx[0]
    me.exch(0, n)
    me.sink(0, n)
    me.size = n
    return ret
}

// O(lgN). Change the priority of the given element.
func (me *PriorityQueue) Change(index int, priority interface{}) {
    if index < 0 || index >= me.capacity {
        panic("Index out of bounds")
    }
    pos := me.pos[index]
    if pos >= me.size {
        panic("Element with that index has already been removed")
    }
    me.a.Change(pos, priority)
    me.swim(pos)
    me.sink(pos, me.size)
}

func (me *PriorityQueue) swim(k int) {
    for {
        i := (k - 1)/2
        if k == 0 || !me.a.Less(i, k) {
            break
        }
        me.exch(i, k)
        k = i
    }
}

func (me *PriorityQueue) sink(k, n int) {
    for {
        i := 1 + (2 * k)
        if i >= n {
            break
        }
        if i + 1 < n && me.a.Less(i, i + 1) {
            i = i + 1
        }
        if me.a.Less(k, i) {
            me.exch(k, i)
            k = i
        } else {
            break
        }
    }
}

func (me *PriorityQueue) exch(i, j int) {
    me.a.Exch(i, j)
    temp := me.idx[i]
    me.idx[i] = me.idx[j]
    me.idx[j] = temp
    me.pos[me.idx[i]] = i
    me.pos[me.idx[j]] = j
}