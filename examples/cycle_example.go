// TODO: Enrich the test suite (since our cycle algorithm is homemade)

package main

import (
    "fmt"
    "strings"
    "github.com/seri/goalgo/graph"
    "./util/graphU"
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
        isDAG := !graph.HasCycle(graphU.ParseGraph(filename))
        if isDAG {
            fmt.Print("Has no cycle")
        } else {
            fmt.Print("Has cycles")
        }
        if isDAG == strings.Contains(filename, "DAG")  {
            fmt.Println(" (Passed)")
        } else {
            fmt.Println(" (Failed)")
        }
    }
}
