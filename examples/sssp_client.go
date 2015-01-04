package main

import (
    "fmt"
    "github.com/seri/goalgo/graph"
    "./util/graphU"
    "./util/reflectU"
)

type Algorithm func(*graph.G, int) *graph.SSSP

type Input struct {
    graf *graph.G
    source, target int
}

type Output struct {
    pathLen int
    dist float64
}

func (me Output) isEqual(you *Output) bool {
    return me.pathLen == you.pathLen && me.dist == you.dist
}

func produceOutput(in *Input, algo Algorithm) *Output {
    sssp := algo(in.graf, in.source)
    return &Output { len(sssp.PathTo(in.target)), sssp.DistTo(in.target) }
}

func printDetail(algo Algorithm, out *Output) {
    fmt.Printf("    %s: DistTo() = %.2f, len(PathTo()) = %d\n",
        reflectU.TypeName(algo), out.dist, out.pathLen)
}

func testAlgoPair(in *Input, a1, a2 Algorithm) bool {
    out1, out2 := produceOutput(in, a1), produceOutput(in, a2)
    if !out1.isEqual(out2) {
        fmt.Println("Failed")
        fmt.Printf("    Source: %d\n", in.source)
        fmt.Printf("    Target: %d\n", in.target)
        printDetail(a1, out1)
        printDetail(a2, out2)
        return false
    }
    return true
}

func testAlgos(in *Input, algos []Algorithm) bool {
    for i := 1; i < len(algos); i++ {
        if !testAlgoPair(in, algos[i], algos[0]) {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println("Test algorithms to find single-source shortest paths")

    filenames := []string {
        "data/tinyEWDAG.txt",
    }

    algos := []Algorithm {
        graph.Dijkstra,
        graph.AcyclicSP,
        graph.BellmanFord,
    }

    for _, filename := range filenames {
        fmt.Println("Input file: " + filename)
        graf := graphU.ParseGraph(filename)
        for source := 0; source < graf.V(); source++ {
            for target := 0; target < graf.V(); target++ {
                in := &Input { graf, source, target }
                if !testAlgos(in, algos) {
                    panic("Bye")
                }
            }
        }
        fmt.Println("OK. All implementations seem to agree with each other.")
    }
}
