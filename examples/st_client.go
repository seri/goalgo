package main

import (
    "fmt"
    "io"
    "math/rand"
    "time"
    "github.com/seri/goalgo/st"
    . "./util"
)

const (
    LogFile = "st_client.log"
)

// Experiment

type Experiment struct {
    size int
    keys []int
    vals []int
    gets []int
    dels []int
}

func NewExperiment(size int) *Experiment {
    keys := make([]int, size)
    vals := make([]int, size)
    gets := make([]int, size)
    dels := make([]int, size)
    for i := 0; i < size; i++ {
        keys[i] = rand.Intn(size)
        vals[i] = rand.Intn(size)
        gets[i] = rand.Intn(size)
        dels[i] = rand.Intn(size)
    }
    return &Experiment { size, keys, vals, gets, dels }
}

func (me Experiment) WriteTo(w io.Writer) {
    fmt.Fprintf(w, "Size: %d\n", me.size)
    fmt.Fprintln(w, "Inserted pairs:")
    for i := 0; i < me.size; i++ {
        fmt.Fprintf(w, "    %d -> %d\n", me.keys[i], me.vals[i])
    }
    fmt.Fprint(w, "Retrieved keys: ")
    for _, x := range me.gets {
        fmt.Fprintf(w, "%d ", x)
    }
    fmt.Fprint(w, "Removed keys: ")
    for _, x := range me.dels {
        fmt.Fprintf(w, "%d ", x)
    }
}

// An ad-hoc symbol table implementation using Go's built-in map

type GoMap map[int]int

func (me GoMap) Empty() bool {
    return len(me) == 0
}

func (me GoMap) Contains(k st.Comparable) bool {
    panic("not supported")
}

func (me GoMap) Get(k st.Comparable) interface{} {
    if v, ok := me[int(k.(st.Int))]; ok {
        return v
    }
    return nil
}

func (me GoMap) Put(k st.Comparable, v interface{}) {
    me[int(k.(st.Int))] = v.(int)
}

func (me GoMap) Flatten() []st.Item {
    panic("not supported")
}

func (me GoMap) Remove(k st.Comparable) {
    delete(me, int(k.(st.Int)))
}

// Benchmark 

func RunPut(xs []st.ST, e *Experiment) {
    for _, x := range xs {
        TimeIt("Inserting into " + Type(x), func () {
            for i := 0; i < e.size; i++ {
                x.Put(st.Int(e.keys[i]), e.vals[i])
            }
        })
    }
}

func RunGet(xs []st.ST, e *Experiment) {
    for _, x := range xs {
        TimeIt("Retrieving from " + Type(x), func() {
            for _, k := range e.gets {
                x.Get(st.Int(k))
            }
        })
    }
}

func RunRemove(xs []st.ST, e *Experiment) {
    for _, x := range xs {
        TimeIt("Removing from " + Type(x), func() {
            for _, k := range e.dels {
                x.Remove(st.Int(k))
            }
        })
    }
}

func RunChecksum(xs []st.ST, e *Experiment) {
    var correct int
    for i, x := range xs {
        TimeIt("Doing checksum in " + Type(x), func() {
            sum := 0
            for _, k := range e.keys {
                v := x.Get(st.Int(k))
                if v != nil {
                    sum = (sum + v.(int)) % Modulus
                }
            }
            switch {
            case i == 0:
                correct = sum
            case sum != correct:
                Fail(LogFile, Type(x) + " is incorrect", e)
            default:
                fmt.Print("(Passed) ")
            }
        })
    }
}

func main() {
    fmt.Println(`
    Test and benchmark symbol-table data structures.
    We expect to see faster retrievals from LLRB compared to BST.
    `)
    rand.Seed(time.Now().UnixNano())
    xs := []st.ST {
        GoMap(make(map[int]int)),
        st.NewBST(),
        st.NewLLRB(),
    }
    es := []*Experiment {
        NewExperiment(TenPow(1)),
        NewExperiment(TenPow(7)),
    }
    for _, e := range es {
        s := PPInt(e.size)
        fmt.Printf("With %s insertions, %s removals, %s retrievals \n", s, s, s)
        RunPut(xs, e)
        RunGet(xs, e)
        RunRemove(xs, e)
        RunChecksum(xs, e)
        fmt.Println()
    }
}