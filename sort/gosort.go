package sort

import (
    gosort "sort"
)

type GoSortable struct {
    Sortable
}

func (me GoSortable) Len() int {
    return me.Sortable.Size()
}

func (me GoSortable) Swap(i, j int) {
    me.Sortable.Exch(i, j)
}

// Sort using Go's standard library.
func GoSort(a Sortable) {
    gosort.Sort( &GoSortable { a } )
}
