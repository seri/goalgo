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
