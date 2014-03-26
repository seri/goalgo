package main

import (
    "fmt"
    "io"
    "math/rand"
    "time"
    "github.com/seri/goalgo/sort"
    "github.com/seri/goalgo/pq"
    . "./util"
)

const (
    LogFile = "sort_client.log"
)

// Experiment

type Experiment struct {
    size int
    max int
    input []int
    output []int
    result []int
}

func NewExperiment(size, max int) *Experiment {
    input := make([]int, size)
    output := make([]int, size)
    result := make([]int, size)
    for i := 0; i < size; i++ {
        input[i] = rand.Intn(max)
        output[i] = input[i]
    }
    return &Experiment { size, max, input, output, result }
}

func (me Experiment) WriteTo(w io.Writer) {
    fmt.Fprintf(w, "Size   : %d\n", me.size)
    fmt.Fprintf(w, "Max    : %d\n", me.max)
    fmt.Fprintf(w, "Input  : %v\n", me.input)
    fmt.Fprintf(w, "Ouput  : %v\n", me.output)
    fmt.Fprintf(w, "Result : %v\n", me.result)
}

func (me Experiment) Reset() {
    copy(me.output, me.input)
}

// Worker

type Worker struct {
    sorter sort.Sorter
    cheater sort.Cheater
    limit int
}

func NewSorter(sorter sort.Sorter, limit int) *Worker {
    return &Worker { sorter, nil, limit }
}

func NewCheater(cheater sort.Cheater, limit int) *Worker {
    return &Worker { nil, cheater, limit }
}

func (me Worker) Name() string {
    if me.sorter != nil {
        return Type(me.sorter)
    }
    return Type(me.cheater)
}

func (me Worker) Sort(a []int) {
    if me.sorter != nil {
        me.sorter(sort.IntSlice(a))
    } else {
        me.cheater(a)
    }
}

// Test

func Check(w *Worker, e *Experiment) {
    for i := range e.input {
        if e.output[i] != e.result[i] {
            fmt.Println()
            Fail(LogFile, w.Name() + " is incorrect", e)
        }
    }
    fmt.Print(" (Passed)")
}

// Benchmark

func RunWorker(w *Worker, e *Experiment) {
    fmt.Printf("Running %-20s .. ", w.Name())

    if e.size > w.limit {
        fmt.Println("Skipped")
        return
    }

    e.Reset()
    t := NewTimer()
    w.Sort(e.output)
    fmt.Printf("%-15s", t.Elapsed())

    if w.Name() == "GoSort" {
        copy(e.result, e.output)
    } else {
        Check(w, e)
    }

    fmt.Println()
}

func RunWorkers(ws []*Worker, e *Experiment) {
    fmt.Printf("With input size %s and maximum value %s\n", PPInt(e.size), PPInt(e.max))
    for _, w := range ws {
        RunWorker(w, e)
    }
    fmt.Println()
}

func RunExperiments(ws []*Worker, es []*Experiment) {
    for _, e := range es {
        RunWorkers(ws, e)
    }
}

func main() {
    fmt.Println(`
    Test and benchmark sorting algorithms.
    
    - When the maximum value is less than the input size, there should be 
    duplicates. We expect to see Djikstra's three-way quicksort perform better 
    than standard quicksort in such cases.
    
    - Mergesort cheats the benchmark by operating directly on []int. Other
    sorters have to go through an interface and suffers a significant overhead.
    `)
    rand.Seed(time.Now().UnixNano())
    ws := []*Worker {
        NewSorter(sort.GoSort, TenPow(7)),
        NewSorter(sort.SelectionSort, TenPow(4)),
        NewSorter(sort.InsertionSort, TenPow(4)),
        NewSorter(sort.ShellSort, TenPow(6)),
        NewCheater(sort.MergeSort, TenPow(7)),
        NewSorter(sort.QuickSort, TenPow(7)),
        NewSorter(sort.DjikstraQuickSort, TenPow(7)),
        NewSorter(pq.HeapSort, TenPow(7)),
    }
    es := []*Experiment {
        NewExperiment(TenPow(1), TenPow(1)),
        NewExperiment(TenPow(4), TenPow(4)),
        NewExperiment(TenPow(6), TenPow(6)),
        NewExperiment(TenPow(7), TenPow(7)),
        NewExperiment(TenPow(7), TenPow(2)),
    }
    RunExperiments(ws, es)
}
