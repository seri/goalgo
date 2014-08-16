package sort

func SelectionSort(a Sortable) {
    for i := 0; i < a.Size() - 1; i++ {
        min := i
        for j := i + 1; j < a.Size(); j++ {
            if a.Less(j, min) {
                min = j
            }
        }
        a.Exch(i, min)
    }
}
