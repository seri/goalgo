// QuickFind is an implementation of the UnionFind API that supports `find` in
// constant time but it suffers linear complexity with regard to `union`. As it
// tries doing more work as early as possible, this is so called the eager 
// approach to the dynamic connectivity problem.

package uf

type QuickFind struct {
    root []int
    count int
}

func NewQuickFind(size int) *QuickFind {
    me := new(QuickFind)
    me.Reset(size)
    return me
}

// ~ N
func (me *QuickFind) Reset(size int) {
    me.root = make([]int, size, size)
    me.count = size
    for i := range(me.root) {
        me.root[i] = i
    }
}

// ~ 1
func (me *QuickFind) Size() int {
    return len(me.root)
}

// ~ 1
func (me *QuickFind) Count() int {
    return me.count
}

// ~ 1
func (me *QuickFind) Connected(p, q int) bool {
    return me.root[p] == me.root[q]
}

// ~ N
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
