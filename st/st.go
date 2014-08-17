// Package st provides the symbol table API. Symbol tables are more commonly known
// as dictionaries or maps. We provide two trees that implement this API.
// 
// Examples:
// 
//		* https://github.com/seri/goalgo/blob/master/examples/st_client.go
//		test and benchmarks the two trees.
// 
// References:
// 
//		1. http://algs4.cs.princeton.edu/32bst/
//		2. http://algs4.cs.princeton.edu/33balanced/
//		3. http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
//		4. http://www.read.seas.harvard.edu/~kohler/notes/llrb.html
package st

// In order to use each ymbol tables, you must first adapt your key type to the
// Comparable interface. 
type Comparable interface {
    Compare(x Comparable) int	// -1 for less, 0 for equal, 1 for greater
}

// When you flatten a symbol table, you will get a slice of Key-value pairs.
type Item struct {
    Key Comparable
    Value interface{}
}

// As you can see, when you retrieve value out of a symbol table with Get(), the
// type of the value is just an empty interface, so you will have to manually
// typecast it.
type ST interface {
    Empty() bool				     // is the symbol table empty
    Contains(k Comparable) bool		 // does it contain the given key
    Get(k Comparable) interface{}    // get the value associated with given key
    Put(k Comparable, v interface{}) // insert a new key-value pair into the symbol table
    Flatten() []Item                 // retrieve all key-value pairs at once
    Remove(k Comparable)             // remove the key-value pair associated with given key
}

// For convenience, we provide an adaptation of int against the Comparable
// interface so you are ready to use int as a key type.
type Int int

func (me Int) Compare(you Comparable) int {
    x, y := int(me), int(you.(Int))
    switch {
    case x < y : return -1
    case x > y : return 1
    }
    return 0
}

// For convenience, we provide an adaptation of string against the Comparable
// interface so you are ready to use string as a key type.
type String string

func (me String) Compare(you Comparable) int {
    x, y := string(me), string(you.(String))
    switch {
    case x < y : return -1
    case x > y : return 1
    }
    return 0
}
