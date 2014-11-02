package main

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    . "./util"
)

func main() {
    filenames := []string {
        "data/tinyDG.txt",
        "data/tinyDAG.txt",
        "data/mediumDG.txt",
        "data/tinyEWDAG.txt",
        "data/tinyEWG.txt",
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
