package graphU

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    "os"
)

func parseGraph(filename string, isDirected bool) *graph.G {
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    var V, E int
    fmt.Fscanln(f, &V)
    fmt.Fscanln(f, &E)

    g := graph.New(V)
    for i := 0; i < E; i++ {
        var (
            u, v int
            weight float64
        )
        fmt.Fscanln(f, &u, &v, &weight)
        g.Add(u, v, weight)
        if !isDirected {
            g.Add(v, u, weight)
        }
    }
    return g
}

func ParseGraph(filename string) *graph.G {
    return parseGraph(filename, true)
}

func ParseUndirectedGraph(filename string) *graph.G {
    return parseGraph(filename, false)
}
