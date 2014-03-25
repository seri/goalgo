package util

import (
    "strconv"
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