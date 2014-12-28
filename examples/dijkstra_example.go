// TODO: Add a check function

package main

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    "./util/graphU"
)

func main() {
    filenames := []string {
        "data/tinyEWD.txt",
        "data/mediumEWD.txt",
    }
	var source = 0
    for _, filename := range filenames {
        fmt.Println("Running Dijkstra's shortest paths against " + filename)
        graf := graphU.ParseGraph(filename)
        sp := graph.Dijkstra(graf, source)
        for target := 0; target < graf.V(); target++ {
			fmt.Printf("%d to %d (%.2f) %v\n",
					   source, target, sp.DistTo(target), sp.PathTo(target))
        }
    }
}
