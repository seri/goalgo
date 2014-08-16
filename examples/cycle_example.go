package main

import (
    "fmt"
    "os"
    "github.com/seri/goalgo/graph"
)

func ParseGraph(filename string) *graph.G {
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    var V, E, u, v int

    fmt.Fscanf(f, "%d %d", &V, &E)
    g := graph.New(V)

    for i := 0; i < E; i++ {
        fmt.Fscanf(f, "%d %d", &u, &v)
        g.Add(u, v, 1)
    }

    return g
}

func main() {
    filenames := []string {
        "data/tinyDG.txt",
        "data/tinyDAG.txt",
    }
    for _, filename := range filenames {
        fmt.Print(filename + ": ")
        if graph.HasCycle(ParseGraph(filename)) {
            fmt.Println("Has cycle")
        } else {
            fmt.Println("Has no cycle")
        }
    }
}
