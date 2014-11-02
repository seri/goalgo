package numberU

import (
    "strconv"
    "fmt"
    "strings"
)

const (
    Modulus = 999331
)

func PPInt(x int) string {
    if x <= 1000 {
        return strconv.Itoa(x)
    }
    n := 0
    for ; x > 1; x /= 10 {
        n++
    }
    return "10^" + strconv.Itoa(n)
}

func TenPow(n int) int {
    ret := 1
    for i := 0; i < n; i++ {
        ret *= 10
    }
    return ret
}

func MinInt(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func ToIntSlice(s string) []int {
    a := make([]int, 0)
    r := strings.NewReader(s)
    for {
        var i int
        n, err := fmt.Fscanf(r, "%d", &i)
        if n != 1 || err != nil {
            break
        }
        a = append(a, i)
    }
    return a
}
