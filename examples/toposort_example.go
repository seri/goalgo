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

func ParseGraph(filename string) (g *graph.G, names []string) {
    var ok bool
    ids := make(map[string]int)
    V := 0

    EachLine(filename, func (lineNo int, line string) {
        jobs := strings.Split(line, "/")
        for _, job := range jobs {
            if _, ok = ids[job]; !ok {
                ids[job] = V
                V++
            }
        }
    })

    names = make([]string, V)
    for job, id := range ids {
        names[id] = job
    }

    g = graph.New(V)

    EachLine(filename, func (lineNo int, line string) {
        jobs := strings.Split(line, "/")
        u := ids[jobs[0]]
        for i := 1; i < len(jobs); i++ {
            v := ids[jobs[i]]
            g.Add(u, v, 1)
        }
    })

    return
}

func main() {
    fmt.Println(`
    A simple example of using graph.TopoSort() to topologically sort the
    directed graph described in data/jobs.txt.
    `)
    g, names := ParseGraph(DataFile)
    ids := graph.TopoSort(g)
    for _, id := range ids {
        fmt.Printf("        %s\n", names[id])
    }
}
