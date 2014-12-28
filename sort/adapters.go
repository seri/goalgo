package sort

// For convenience, we provide an adaptation of []int against the Sortable
// interface.
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

// For convenience, we provide an adaptation of []string against the Sortable
// interface.
type StringSlice []string

func (me StringSlice) Size() int {
    return len(me)
}
func (me StringSlice) Less(i, j int) bool {
    return me[i] < me[j]
}
func (me StringSlice) Exch(i, j int) {
    me[i], me[j] = me[j], me[i]
}

// This trick allows one to quickly reverse the order of a Sortable
type rSortable struct {
    Sortable
}
func (me rSortable) Less(i, j int) bool {
    return me.Sortable.Less(j, i)
}

// Tweak a Sortable collection so that when a Sorter is used, the order will be
// reversed. Which means you can sort a collection in descending order like this:
//      sort.QuickSort(sort.Reverse(collection))
func Reverse(x Sortable) Sortable {
    return &rSortable { x }
}
