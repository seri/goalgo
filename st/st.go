/*
    The st package provides the symbol table API, which is more commonly known
    as dictionaries or maps, and two trees that implement it.

    In order to use each of these symbol tables, you must first adapt your key
    type to the Comparable interface. When you get values out of a ST, you also
    have to manually type-cast them. Have a look at the runnable st_client.go
    for details.

    See:
    1) http://algs4.cs.princeton.edu/32bst/
    2) http://algs4.cs.princeton.edu/33balanced/
    3) http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
    4) http://www.read.seas.harvard.edu/~kohler/notes/llrb.html
*/

package st

type Comparable interface {
    Compare(x Comparable) int
}

type Item struct {
    Key Comparable
    Value interface{}
}

type ST interface {
    Empty() bool
    Contains(k Comparable) bool
    Get(k Comparable) interface{}
    Put(k Comparable, v interface{})
    Flatten() []Item
    Remove(k Comparable)
}

// Common key types

type Int int

func (me Int) Compare(you Comparable) int {
    x, y := int(me), int(you.(Int))
    switch {
    case x < y : return -1
    case x > y : return 1
    }
    return 0
}

type String string

func (me String) Compare(you Comparable) int {
    x, y := string(me), string(you.(String))
    switch {
    case x < y : return -1
    case x > y : return 1
    }
    return 0
}
