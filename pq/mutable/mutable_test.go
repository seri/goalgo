package mutable

import (
    "fmt"
    "testing"
    "github.com/seri/goalgo/sort"
    "github.com/seri/goalgo"
)

const (
    N = 100
)

func Example() {
    a := []int {0, 3, 2, 1} // 0 -> 0, 1 -> 3, 2 -> 2, 3 -> 1
    queue := New(IntSlice(a))

    if queue.Size() != 4 {
        fmt.Printf("Size() returns %d (expecting 4)", queue.Size())
    }

    index := queue.Pop()
    if index != 1 {
        fmt.Printf("Pop() returns %d (expecting 1)", index)
    }

    queue.Change(0, 4)
    index = queue.Pop()
    if index != 1 {
        fmt.Printf("Pop() returns %d (expecting 0)", index)
    }
}

func prepare() (a, b []int) {
    a = make([]int, N)
    b = make([]int, N)
    for i := range a {
        a[i] = i
    }
    goalgo.Shuffle(sort.IntSlice(a))
    for i := range a {
        b[a[i]] = i
    } 
    return
}

func TestPop(t *testing.T) {
    a, b := prepare()
    pq := New(IntSlice(a))
    for i := N - 1; i >= 0; i-- {
        if pq.Pop() != b[i] {
            t.FailNow()            
        }
    }    
}

func TestChange(t *testing.T) {
    a, b := prepare()
    pq := New(IntSlice(a))
    goalgo.Shuffle(sort.IntSlice(b))
    for _, index := range b {
        pq.Change(index, N)
        if pq.Pop() != index {
            t.Fail()            
        }
    }    
}

func TestReverse(t *testing.T) {
    a, b := prepare()
    pq := New(Reverse(IntSlice(a)))
    for i := range b {
        if pq.Pop() != b[i] {
            t.FailNow()            
        }
    }    
}