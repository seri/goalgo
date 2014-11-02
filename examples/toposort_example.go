package main

import (
    "fmt"
    "strings"
    "github.com/seri/goalgo/graph"
    . "./util"
)

const (
    DataFile = "data/jobs.txt"
)

func CollectJobs(filename string) map[string]int {
    table := make(map[string]int)
    n := 0
    EachLine(filename, func (lineNo int, line string) {
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

func ParseGraph(filename string) (g *graph.G, jobs []string) {
    table := CollectJobs(filename)
    V := len(table)
    jobs = make([]string, V)
    for job, i := range table {
        jobs[i] = job
    }
    g = graph.New(V)
    EachLine(filename, func (lineNo int, line string) {
        parts := strings.Split(line, "/")
        u := table[parts[0]]
        for i := 1; i < len(parts); i++ {
            v := table[parts[i]]
            g.Add(u, v, 1)
        }
    })
    return
}

func SanityCheck(g *graph.G, topoOrder []int) {
    indices := make([]int, g.V())
    for i, u := range topoOrder {
        indices[u] = i
    }
    for i, u := range topoOrder {
        for _, v := range g.Vertices(u) {
            if indices[v] < i {
                panic("Invalid topological order")
            }
        }
    }
    fmt.Println("The order seems to be valid")
}

func main() {
    fmt.Println(`A simple example of using graph.TopoSort() to topologically sort
the directed graph described in data/jobs.txt.`)
    g, jobs := ParseGraph(DataFile)
    topoOrder := graph.TopoSort(g)
    for _, u := range topoOrder {
        fmt.Printf("    %s\n", jobs[u])
    }
    SanityCheck(g, topoOrder)
}
