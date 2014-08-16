package pq

import (
    "github.com/seri/goalgo/sort"
)

type adapter struct {
    a sort.Sortable
    n int
}

func (me *adapter) Size() int {
    return me.n
}

func (me *adapter) Less(i, j int) bool {
    return me.a.Less(i, j)
}

func (me *adapter) Exch(i, j int) {
    me.a.Exch(i, j)
}

func (me *adapter) Push(x interface{}) {
    panic("not supported")
}

func (me *adapter) Pop() interface{} {
    me.n--
    return nil
}

func HeapSort(a sort.Sortable) {
    b := adapter { a, a.Size() }
    Heapify(&b)
    for b.n > 0 {
        Pop(&b)
    }
}
