package pq

import (
    "github.com/seri/goalgo/sort"
)

type Interface interface {
    sort.Sortable
    Push(x interface{})
    Pop() interface{}
}

// ~ N
func Heapify(a Interface) {
    n := a.Size()
    for i := n/2 - 1; i >= 0; i-- {
        sink(a, i, n)
    }
}

// ~ lgN
func Push(a Interface, x interface{}) {
    a.Push(x)
    swim(a, a.Size() - 1)
}

// ~ lgN
func Pop(a Interface) interface{} {
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
