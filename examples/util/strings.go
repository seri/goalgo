package util

import (
	"fmt"
	"strings"
)

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