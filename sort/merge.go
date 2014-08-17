package sort

import (
    "math"
)

func merge(a, x, y []int) {
    i, j, k := 0, 0, 0
    for {
        if x[i] <= y[j] {
            a[k] = x[i]
            k, i = k + 1, i + 1
            if i == len(x) {
                copy(a[k:len(a)], y[j:len(y)])
                return
            }
        } else {
            a[k] = y[j]
            k, j = k + 1, j + 1
            if j == len(y) {
                copy(a[k:len(a)], x[i:len(x)])
                return
            }
        }
    }
}

func mergeSort(a, b []int) {
    n := len(a)
    if n < 2 {
        return
    }
    mid := int(math.Ceil(float64(n)/2))
    mergeSort(b[0:mid], a[0:mid]) // b[0:mid] is now sorted
    mergeSort(b[mid:n], a[mid:n]) // b[mid:n] is now sorted
    merge(a, b[0:mid], b[mid:n])  // merge them into a so a is now sorted
}

// Merge sort is theoretically optimal with guaranteed ~NlgN compares, although in
// practice it is outperformed by quick sort. Depending on the merging strategy, merge
// sort may be stable or unstable. Our implementation produces a stable sort. Notice
// that merge sort only accepts an int slice. This is because our implementation
// depends on creating a copy of the initial collection which cannot be done in Go in
// a generic manner for the moment.
func MergeSort(a []int) {
    b := make([]int, len(a))
    copy(b, a)
    mergeSort(a, b)
}
