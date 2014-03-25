// Sort with Go's standard library. This is used to check the correctness and
// performances of other sorters.

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

func GoSort(a Sortable) {
    gosort.Sort( &GoSortable { a } )
}