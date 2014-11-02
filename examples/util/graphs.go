package util

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    "os"
)

func ParseGraph(filename string) *graph.G {
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
            weight float32
        )
        fmt.Fscanln(f, &u, &v, &weight)
        g.Add(u, v, 1)
    }
    return g
}

