package main

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    "./util/graphU"
)

func main() {
    filenames := []string {
        "data/tinyDG.txt",
        "data/mediumDG.txt",
    }
    for _, filename := range filenames {
        fmt.Print("Running SCC against " + filename + " ... ")
        scc := graph.SCC(graphU.ParseGraph(filename))
        fmt.Printf("%d components\n", len(scc))
        for _, component := range scc {
            fmt.Println(component)
        }
    }
}
