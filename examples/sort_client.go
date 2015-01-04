package main

import (
    "fmt"
    "io"
    "math/rand"
    "time"
    "github.com/seri/goalgo/sort"
    "github.com/seri/goalgo/pq"
    "./util/numberU"
    "./util/logU"
    "./util/reflectU"
    "./util/timer"
)

const (
    LogFile = "sort_client.log"
)

// The experiment data type encapsulating a sorting test case

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

// The worker data type which wraps a sorter with some added info

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
        return reflectU.TypeName(me.sorter)
    }
    return reflectU.TypeName(me.cheater)
}

func (me Worker) Sort(a []int) {
    if me.sorter != nil {
        me.sorter(sort.IntSlice(a))
    } else {
        me.cheater(a)
    }
}

// Run workers against experiments

func check(w *Worker, e *Experiment) {
    for i := range e.input {
        if e.output[i] != e.result[i] {
            fmt.Println()
            logU.Fail(LogFile, w.Name() + " is incorrect", e)
        }
    }
    fmt.Print(" (Passed)")
}

func runWorker(w *Worker, e *Experiment) {
    fmt.Printf("Running %-20s .. ", w.Name())

    if e.size > w.limit {
        fmt.Println("Skipped")
        return
    }

    e.Reset()
    t := timer.New()
    w.Sort(e.output)
    fmt.Printf("%-15s", t.Elapsed())

    if w.Name() == "GoSort" {
        copy(e.result, e.output)
    } else {
        check(w, e)
    }

    fmt.Println()
}

func runWorkers(ws []*Worker, e *Experiment) {
    fmt.Printf("With input size %s and maximum value %s\n",
               numberU.PPInt(e.size), numberU.PPInt(e.max))
    for _, w := range ws {
        runWorker(w, e)
    }
    fmt.Println()
}

func runExperiments(ws []*Worker, es []*Experiment) {
    for _, e := range es {
        runWorkers(ws, e)
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
        NewSorter(sort.GoSort,              numberU.TenPow(7)),
        NewSorter(sort.SelectionSort,       numberU.TenPow(4)),
        NewSorter(sort.InsertionSort,       numberU.TenPow(4)),
        NewSorter(sort.ShellSort,           numberU.TenPow(6)),
        NewCheater(sort.MergeSort,          numberU.TenPow(7)),
        NewSorter(sort.QuickSort,           numberU.TenPow(7)),
        NewSorter(sort.DjikstraQuickSort,   numberU.TenPow(7)),
        NewSorter(pq.HeapSort,              numberU.TenPow(7)),
    }
    es := []*Experiment {
        NewExperiment(numberU.TenPow(1), numberU.TenPow(1)),
        NewExperiment(numberU.TenPow(4), numberU.TenPow(4)),
        NewExperiment(numberU.TenPow(6), numberU.TenPow(6)),
        NewExperiment(numberU.TenPow(7), numberU.TenPow(7)),
        NewExperiment(numberU.TenPow(7), numberU.TenPow(2)),
    }
    runExperiments(ws, es)
}
