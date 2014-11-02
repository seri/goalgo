package main

import (
    "strings"
    "./util/testU"
)

func StandardPartition(s string) string {
    a := strings.Split(s, " ")
    x, i, j := a[0], 1, len(a) - 1
    for {
        for ; i < len(a) && a[i] < x; i++ { }
        for ; a[j] > x; j-- { }
        if i >= j {
            break
        }
        a[i], a[j] = a[j], a[i]
        i++
        j--
    }
    a[0], a[j] = a[j], a[0]
    return strings.Join(a, " ")
}

func DjkstraPartition(s string) string {
    a := strings.Split(s, " ")
    x, lt, gt := a[0], 0, len(a) - 1
    for i := lt; i <= gt; {
        switch {
        case a[i] < x:
            a[i], a[lt] = a[lt], a[i]
            lt++
            i++
        case a[i] > x:
            a[i], a[gt] = a[gt], a[i]
            gt--
        default:
            i++
        }
    }
    return strings.Join(a, " ")
}

func main() {
    testU.ExpectOutput(StandardPartition,
        "49 97 84 93 25 33 72 40 54 96 19 30",
        "33 30 19 40 25 49 72 93 54 96 84 97")
    testU.ExpectOutput(DjkstraPartition,
         "53 53 93 55 87 65 91 29 27 53",
         "27 29 53 53 53 91 65 87 55 93")
    testU.GetOutput(StandardPartition,
          "A A B B B B A B A B A B")
    testU.GetOutput(DjkstraPartition,
          "43 55 43 43 44 25 22 93 50 66")
}
