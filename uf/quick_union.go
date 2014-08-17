package uf

type QuickUnion struct {
    parent []int
    size []int
    count int
}

// Create a new quick union instance with the given number of vertices.
func NewQuickUnion(size int) *QuickUnion {
    me := new(QuickUnion)
    me.Reset(size)
    return me
}

// Reset this quick union instance to its initial state.
func (me *QuickUnion) Reset(size int) {
    me.parent = make([]int, size, size)
    me.size = make([]int, size, size)
    me.count = size
    for i := range me.parent {
        me.parent[i] = i
        me.size[i] = 1
    }
}

// O(1). Retrieve the number of vertices.
func (me *QuickUnion) Size() int {
    return len(me.parent)
}

// O(1). Retrieve the number of connected components.
func (me *QuickUnion) Count() int {
    return me.count
}

// ~ lgN
func (me *QuickUnion) root(p int) int {
    for p != me.parent[p] {
        p = me.parent[p]
    }
    return p
}

// O(lgN). Whether there is a path between the two vertices p and q.
func (me *QuickUnion) Connected(p, q int) bool {
    return me.root(p) == me.root(q)
}

// ~ 1
func (me *QuickUnion) plug(parent, child int) {
    me.parent[child] = parent
    me.size[parent] += me.size[child]
}

// O(lgN). Add an edge between p and q.
func (me *QuickUnion) Union(p, q int) {
    pp, qq := me.root(p), me.root(q)
    if pp != qq {
        if (me.size[pp] < me.size[qq]) {
            me.plug(qq, pp)
        } else {
            me.plug(pp, qq)
        }
        me.count -= 1
    }
}
