// QuickUnion is an implementation of the UnionFind API that supports both 
// `find` and `union` in lgN time. As it delays doing work until the time
// comes, it's a lazy approach to the dynamic connectivity problem.

package uf

type QuickUnion struct {
    parent []int
    size []int
    count int
}

func NewQuickUnion(size int) *QuickUnion {
    me := new(QuickUnion)
    me.Reset(size)
    return me
}

// ~ 1
func (me *QuickUnion) Reset(size int) {
    me.parent = make([]int, size, size)
    me.size = make([]int, size, size)
    me.count = size
    for i := range me.parent {
        me.parent[i] = i
        me.size[i] = 1
    }
}

// ~ 1
func (me *QuickUnion) Size() int {
    return len(me.parent)
}

// ~ 1
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

// ~ lgN
func (me *QuickUnion) Connected(p, q int) bool {
    return me.root(p) == me.root(q)
}

// ~ 1
func (me *QuickUnion) plug(parent, child int) {
    me.parent[child] = parent
    me.size[parent] += me.size[child]
}

// ~ lgN
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
