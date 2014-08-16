package sort

func ShellSort(a Sortable) {
    var step int
    for step = 1; step < a.Size(); step = step*3 + 1 {}
    for ; step > 0; step = (step - 1)/3 {
        insertionSort(a, step)
    }
}
