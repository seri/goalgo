package uf

type QuickFind struct {
    root []int
    count int
}

// Create a new quick find instance with the given number of vertices.
func NewQuickFind(size int) *QuickFind {
    me := new(QuickFind)
    me.Reset(size)
    return me
}

// Reset this quick find instance to its initial state.
func (me *QuickFind) Reset(size int) {
    me.root = make([]int, size, size)
    me.count = size
    for i := range(me.root) {
        me.root[i] = i
    }
}

// O(1). Retrieve the number of vertices.
func (me *QuickFind) Size() int {
    return len(me.root)
}

// O(1). Retrieve the number of connected components.
func (me *QuickFind) Count() int {
    return me.count
}

// O(1). Whether there is a path between the two vertices p and q.
func (me *QuickFind) Connected(p, q int) bool {
    return me.root[p] == me.root[q]
}

// O(N). Add an edge between p and q.
func (me *QuickFind) Union(p, q int) {
    pp, qq := me.root[p], me.root[q]
    if pp != qq {
        for i := range(me.root) {
            if me.root[i] == pp {
                me.root[i] = qq
            }
        }
        me.count--
    }
}
