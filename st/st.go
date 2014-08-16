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
