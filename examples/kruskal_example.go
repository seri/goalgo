// TODO: Add a check function

package main

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    "./util/graphU"
)

func main() {
    filenames := []string {
        "data/tinyEWG.txt",
        "data/mediumEWG.txt",
    }
    for _, filename := range filenames {
        fmt.Println("Constructing MST against " + filename)
        graf := graphU.ParseUndirectedGraph(filename)
        mst := graph.Kruskal(graf)
        edges := graph.AllEdges(mst)
        var sum float64 = 0
        for _, edge := range edges {
            if len(edges) < 20 {
                fmt.Printf("    %v\n", edge)
            }
            sum += edge.Weight()
        }
        fmt.Printf("Total cost: %f\n", sum)
    }
}
