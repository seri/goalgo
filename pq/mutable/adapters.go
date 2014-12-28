package mutable

// Adapt an int slice to satisfy mutable.Interface so it can behave as the
// container for a mutable priority queue where priorities are of type int.
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

func (me IntSlice) Change(i int, x interface{}) {
    me[i] = x.(int)
}

// Adapt a float64 slice to satisfy mutable.Interface so it can behave as the
// container for a priority queue where priorities are of type float64.
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

func (me Float64Slice) Change(i int, x interface{}) {
    me[i] = x.(float64)
}

// We only implemenet a maximum priority queue, but if you want a minimum
// priority queue, you can just call New(Reverse(IntSlice(a))).
func Reverse(x Interface) Interface {
    return &rInterface { x }
}

type rInterface struct {
    Interface
}
func (me rInterface) Less(i, j int) bool {
    return me.Interface.Less(j, i)
}
