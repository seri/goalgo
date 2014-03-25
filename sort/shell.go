// ShellSort is an interesting generalisation of InsertionSort whose 
// performance is unknown. This particular implementation of ShellSort uses
// Knuth's (3n + 1) gap sequence, which achieves O(N^(3/2)) complexity.

package sort

func ShellSort(a Sortable) {
    var step int
    for step = 1; step < a.Size(); step = step*3 + 1 {}
    for ; step > 0; step = (step - 1)/3 {
        insertionSort(a, step)
    }
}