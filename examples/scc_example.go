package main

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    "./util/graphU"
)

func check(graf *graph.G, scc [][]int) {
    fmt.Print("Checking the result ... ")
    inSame := graph.InSameComponent(graf, scc)
    reachable := graph.Reachable(graf)
    for u := 0; u < graf.V(); u++ {
        for v := 0; v < graf.V(); v++ {
            if inSame[u][v] != (reachable[u][v] && reachable[v][u]) {
                fmt.Println("Failed")
                fmt.Printf("    inSame[%d][%d]    = %v\n", u, v, inSame[u][v])
                fmt.Printf("    reachable[%d][%d] = %v\n", u, v, reachable[u][v])
                fmt.Printf("    reachable[%d][%d] = %v\n", v, u, reachable[v][u])
                panic("Bye")
            }
        }
    }
    fmt.Println("Passed")
}

func main() {
    filenames := []string {
        "data/tinyDG.txt",
        "data/mediumDG.txt",
    }
    for _, filename := range filenames {
        fmt.Print("Running SCC against " + filename + " ... ")
        graf := graphU.ParseGraph(filename)
        scc := graph.StrongComponents(graf)
        fmt.Printf("%d components\n", len(scc))
        for _, component := range scc {
            fmt.Println(component)
        }
        check(graf, scc)
    }
}
