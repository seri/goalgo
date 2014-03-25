// InsertionSort is a quadratic sorter. It is stable, fast for small input, and 
// performs well in partially sorted collections. The last point is especially
// important as this is the reason why InsertSort is the only quadratic sorter 
// that is still widely used in practice. It uses sqr(N)/4 calls to both Less()
// and Exch() on average.

package sort

func insertionSort(a Sortable, step int) {
    for i := step; i < a.Size(); i++ {
        for j := i; j >= step && a.Less(j, j - step); j -= step {
            a.Exch(j, j - step)
        }
    }
}

func InsertionSort(a Sortable) {
    insertionSort(a, 1)
}