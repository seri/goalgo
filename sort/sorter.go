package sort

type Sortable interface {
    Size() int
    Less(i, j int) bool
    Exch(i, j int)
}

// A Sorter is a function that sorts a Sortable
type Sorter func(Sortable)

// A Cheater is very lame as it can only sort an int slice
type Cheater func([]int)


// Some popular sortable types for convenience's sake

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

func Reverse(x Sortable) Sortable {
    return &rSortable { x }
}
