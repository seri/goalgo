package main

import (
    "fmt"
    "strings"
    "github.com/seri/goalgo/pq"
    "./util/testU"
    "./util/numberU"
)

// Adapt an int slice to satisfy pq.Interface so it can behave as a heap container

type IntSlice []int

func (me IntSlice) Size() int {
    return len(me)
}

func (me IntSlice) Less(i, j int) bool {
    return me[i] < me[j]
}

func (me IntSlice) Exch(i, j int) {
    me[i], me[j] = me[j], me[i]
}

func (me *IntSlice) Push(x interface{}) {
    *me = append(*me, x.(int))
}

func (me *IntSlice) Pop() interface{} {
    a := *me
    n := len(a) - 1
    x := a[n]
    *me = a[:n]
    return x
}

func (me IntSlice) String() string {
    s := fmt.Sprint([]int(me))
    return s[1:len(s) - 1]
}

// Solve the exercises

func HeapInsert(s string) string {
    a := strings.Split(s, ";")
    b, c := IntSlice(numberU.ToIntSlice(a[0])), numberU.ToIntSlice(a[1])
    for _, x := range c {
        pq.Push(&b, x)
    }
    return b.String()
}

func HeapRemove3Times(s string) string {
    a := IntSlice(numberU.ToIntSlice(s))
    for i := 0; i < 3; i++ {
        pq.Pop(&a)
    }
    return a.String()
}

func main() {
    testU.ExpectOutput(HeapInsert,
         "84 78 72 30 40 51 46 21 16 15;98 49 76",
         "98 84 76 30 78 72 46 21 16 15 40 49 51")
    testU.ExpectOutput(HeapRemove3Times,
         "86 71 78 21 67 57 49 20 10 54",
         "67 21 57 20 10 54 49")
    testU.GetOutput(HeapInsert,
          "95 92 36 64 73 31 14 27 48 46;90 24 84")
    testU.GetOutput(HeapRemove3Times,
          "88 87 57 62 75 40 30 26 52 33")
}
