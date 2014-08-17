package sort

func insertionSort(a Sortable, step int) {
    for i := step; i < a.Size(); i++ {
        for j := i; j >= step && a.Less(j, j - step); j -= step {
            a.Exch(j, j - step)
        }
    }
}

// Insertion sort uses sqr(N)/4 compares and exchanges on average. The good news
// is that it is stable and fast in partially sorted collections, which is why
// insertion sort is still idely used in practice.
func InsertionSort(a Sortable) {
    insertionSort(a, 1)
}
