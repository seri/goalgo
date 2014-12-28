package main

import (
    "fmt"
    "strings"
    "github.com/seri/goalgo/pq"
    "./util/testU"
    "./util/numberU"
)

func toString(me pq.IntSlice) string {
    s := fmt.Sprint([]int(me))
    return s[1:len(s) - 1]
}

// Solve the exercises

func heapInsert(s string) string {
    a := strings.Split(s, ";")
    b, c := pq.IntSlice(numberU.ToIntSlice(a[0])), numberU.ToIntSlice(a[1])
    for _, x := range c {
        pq.Push(&b, x)
    }
    return toString(b)
}

func heapRemoveThrice(s string) string {
    a := pq.IntSlice(numberU.ToIntSlice(s))
    for i := 0; i < 3; i++ {
        pq.Pop(&a)
    }
    return toString(a)
}

func main() {
    testU.ExpectOutput(heapInsert,
         "84 78 72 30 40 51 46 21 16 15;98 49 76",
         "98 84 76 30 78 72 46 21 16 15 40 49 51")
    testU.ExpectOutput(heapRemoveThrice,
         "86 71 78 21 67 57 49 20 10 54",
         "67 21 57 20 10 54 49")
    testU.GetOutput(heapInsert,
          "95 92 36 64 73 31 14 27 48 46;90 24 84")
    testU.GetOutput(heapRemoveThrice,
          "88 87 57 62 75 40 30 26 52 33")
}
