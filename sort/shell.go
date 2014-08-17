package sort

// Shell sort is an interesting generalisation of insertion sort whose performance
// is, interestingly enough, still unknown. This particular implementation of
// shell sort uses Knuth's 3n + 1 gap sequence, which achieves O(N^(3/2)) complexity.
func ShellSort(a Sortable) {
    var step int
    for step = 1; step < a.Size(); step = step*3 + 1 {}
    for ; step > 0; step = (step - 1)/3 {
        insertionSort(a, step)
    }
}
