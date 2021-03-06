package main

import (
    "strings"
    "./util/testU"
)

var (
    Count int
)

type Algo func(a, b []string, lo, hi int)

func intMin(p, q int) int {
    if p < q {
        return p
    }
    return q
}

func merge(a, b []string, lo, mid, hi int) {
    if Count == 0 {
        return
    }
    Count--
    for i := lo; i < hi; i++ {
        b[i] = a[i]
    }
    i, j := lo, mid
    for k := lo; k < hi; k++ {
        switch {
        case i >= mid:
            a[k] = b[j]
            j++
        case j >= hi:
            a[k] = b[i]
            i++
        case b[j] < b[i]:
            a[k] = b[j]
            j++
        default:
            a[k] = b[i]
            i++
        }
    }
}

func mergeSort(a, b []string, lo, hi int) {
    if hi <= lo + 1 {
        return
    }
    mid := lo + (hi - lo)/2
    mergeSort(a, b, lo, mid)
    mergeSort(a, b, mid, hi)
    merge(a, b, lo, mid, hi)
}

func buMergeSort(a, b []string, lo, hi int) {
    for n := 2; n <= len(a); n += n {
        for i := 0; i < len(a); i += n {
            merge(a, b, i, i + n/2, intMin(i + n, len(a)))
        }
    }
}

func makeProcessor(algo Algo, n int) testU.Processor {
    Count = n
    return func(s string) string {
        a := strings.Split(s, " ")
        b := make([]string, len(a))
        copy(b, a)
        algo(a, b, 0, len(a))
        return strings.Join(a, " ")
    }
}

func merge7Times(s string) string {
    return makeProcessor(mergeSort, 7)(s)
}

func buMerge7Times(s string) string {
    return makeProcessor(buMergeSort, 7)(s)
}

func main() {
    testU.ExpectOutput(merge7Times,
         "65 82 60 97 93 74 64 52 87 45 32 15",
         "60 65 74 82 93 97 52 64 87 45 32 15")
    testU.ExpectOutput(merge7Times,
         "82 84 59 90 31 52 17 88 24 92 83 91",
         "31 52 59 82 84 90 17 24 88 92 83 91")
    testU.ExpectOutput(buMerge7Times,
         "76 13 38 94 47 74 29 64 46 41",
         "13 38 76 94 29 47 64 74 41 46")
    testU.GetOutput(merge7Times,
          "82 76 58 97 21 88 34 33 16 67 93 64")
}
