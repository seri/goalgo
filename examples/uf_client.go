package main

import (
    "fmt"
    "math/rand"
    "io"
    "time"
    "github.com/seri/goalgo/uf"
    "./util/numberU"
    "./util/reflectU"
    "./util/logU"
    "./util/timer"
)

const (
    LogFile = "uf_client.log"
)

// The experiment data type containing pairs of vertices to connect

type Experiment struct {
    size int
    ps, qs []int
}

func NewExperiment(size int) *Experiment {
    n := rand.Intn(size)
    ps, qs := make([]int, n, n), make([]int, n, n)
    for i := 0; i < n; i++ {
        ps[i] = rand.Intn(size)
        qs[i] = rand.Intn(size)
    }
    return &Experiment { size, ps, qs }
}

func (me Experiment) WriteTo(w io.Writer) {
    fmt.Fprintf(w, "Size: %d\n", me.size)
    fmt.Fprintf(w, "Number of unions: %d\n", len(me.ps))
    fmt.Fprintln(w, "Union details:")
    for i := range me.ps {
        fmt.Fprintf(w, "    %d %d\n", me.ps[i], me.qs[i])
    }
}

// check whether all union find implementations return the same result for
// a particular experiment

func check(ufs []uf.UnionFind, e *Experiment) {
    fmt.Print("Cheking if all implementations match ... ")
    for i := 1; i < len(ufs); i++ {
        if ufs[i].Count() != ufs[0].Count() {
            logU.Fail(LogFile, 
                fmt.Sprintf("%s.Count() returns %d while %s.Count() returns %d",
                reflectU.TypeName(ufs[0]), ufs[0].Count(), 
                reflectU.TypeName(ufs[i]), ufs[i].Count()), e)
        }
    }
    n := e.size
    for i := 0; i < n; i++ {
        p, q := rand.Intn(n), rand.Intn(n)
        x := ufs[0].Connected(p, q)
        for i := 1; i < len(ufs); i++ {
            if ufs[i].Connected(p, q) != x {
                logU.Fail(LogFile, 
                    fmt.Sprintf("%s.Connected(%d, %d) returns %t while " +
                    "%s.Connected(%d, %d) returns %t", reflectU.TypeName(ufs[0]), 
                    p, q, x, reflectU.TypeName(ufs[i]), 
                    p, q, ufs[i].Connected(p, q)), e)
            }
        }
    }
    fmt.Println("Passed")
}

// Run union find implementations against experiments

func runOne(uf uf.UnionFind, e *Experiment) {
    timer.TimeIt("Running " + reflectU.TypeName(uf), func() {
        uf.Reset(e.size)
        for i := range e.ps {
            uf.Union(e.ps[i], e.qs[i])
        }
    })
}

func runAll(ufs []uf.UnionFind, size int) {
    e := NewExperiment(size)
    fmt.Printf("With input size %s and %d calls to Union()\n", numberU.PPInt(e.size), len(e.ps))
    for _, uf := range ufs {
        runOne(uf, e)
    }
    check(ufs, e)
    fmt.Println()
}

func main() {
    rand.Seed(time.Now().UnixNano())
    ufs := []uf.UnionFind { uf.NewQuickFind(10), uf.NewQuickUnion(10) }
    sizes := []int {
        numberU.TenPow(1),
        numberU.TenPow(4),
        numberU.TenPow(5),
    }
    for _, size := range sizes {
        runAll(ufs, size)
    }
}
