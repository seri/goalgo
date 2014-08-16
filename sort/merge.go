// MergeSort is theoretically optimal with guaranteed ~NlgN compares. This non
// in-place MergeSort implementation provided in the course cannot be done with 
// Go in a generic manner. That means MergeSort is not a sort.Sorter.

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

func MergeSort(a []int) {
    b := make([]int, len(a))
    copy(b, a)
    mergeSort(a, b)
}
