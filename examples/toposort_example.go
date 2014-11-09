package main

import (
    "fmt"
    "strings"
    "github.com/seri/goalgo/graph"
    "./util/ioU"
)

const (
    DataFile = "data/jobs.txt"
)

func collectJobs(filename string) map[string]int {
    table := make(map[string]int)
    n := 0
    ioU.EachLine(filename, func (lineNo int, line string) {
        jobs := strings.Split(line, "/")
        for _, job := range jobs {
            if _, ok := table[job]; !ok {
                table[job] = n
                n++
            }
        }
    })
    return table
}

func parseGraph(filename string) (g *graph.G, jobs []string) {
    table := collectJobs(filename)
    V := len(table)
    jobs = make([]string, V)
    for job, i := range table {
        jobs[i] = job
    }
    g = graph.New(V)
    ioU.EachLine(filename, func (lineNo int, line string) {
        parts := strings.Split(line, "/")
        u := table[parts[0]]
        for i := 1; i < len(parts); i++ {
            v := table[parts[i]]
            g.Add(u, v, 1)
        }
    })
    return
}

func check(g *graph.G, topoOrder []int) {
    fmt.Print("Checking the result ... ")
    indices := make([]int, g.V())
    for i, u := range topoOrder {
        indices[u] = i
    }
    for i, u := range topoOrder {
        for _, v := range g.Vertices(u) {
            if indices[v] < i {
                fmt.Println("Failed")
                fmt.Printf("    There is an edge %d->%d\n", u, v)
                fmt.Printf("    But vertice %d appears before %d\n", u, v)
                panic("Bye")
            }
        }
    }
    fmt.Println("Passed")
}

func main() {
    fmt.Println(`A simple example of using graph.TopoSort() to topologically sort
the directed graph described in data/jobs.txt.`)
    g, jobs := parseGraph(DataFile)
    topoOrder := graph.TopoSort(g)
    for _, u := range topoOrder {
        fmt.Printf("    %s\n", jobs[u])
    }
    check(g, topoOrder)
}
