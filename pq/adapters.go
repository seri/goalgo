package pq

// Adapt an int slice to satisfy pq.Interface so it can behave as a priority
// queue where priorities are of type int.
type IntSlice []int

func (me IntSlice) Size() int {
    return len(me)
}

func (me IntSlice) Less(i, j int) bool {
    return me[i] < me[j]
}

func (me IntSlice) Exch(i, j int) {
    me[i], me[j] = me[j], me[i]
}

func (me *IntSlice) Push(x interface{}) {
    *me = append(*me, x.(int))
}

func (me *IntSlice) Pop() interface{} {
    a := *me
    n := len(a) - 1
    x := a[n]
    *me = a[:n]
    return x
}

// Adapt a float64 slice to satisfy pq.Interface so it can behave as a priority
// queue where priorities are of type float64.
type Float64Slice []float64

func (me Float64Slice) Size() int {
    return len(me)
}

func (me Float64Slice) Less(i, j int) bool {
    return me[i] < me[j]
}

func (me Float64Slice) Exch(i, j int) {
    me[i], me[j] = me[j], me[i]
}

func (me *Float64Slice) Push(x interface{}) {
    *me = append(*me, x.(float64))
}

func (me *Float64Slice) Pop() interface{} {
    a := *me
    n := len(a) - 1
    x := a[n]
    *me = a[:n]
    return x
}

// We only implemenet a maximum priority queue, but if you want a minimum
// priority queue, you can just call pq.Heapify(pq.Reverse(pq.IntSlice(a))).
func Reverse(x Interface) Interface {
    return &rInterface { x }
}

type rInterface struct {
    Interface
}
func (me rInterface) Less(i, j int) bool {
    return me.Sortable.Less(j, i)
}
