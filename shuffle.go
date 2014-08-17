// Package goalgo provides implementations of many fundamental algorithms, most
// which come from the Algorithms course of the university of Princeton. There are
// runnable examples for most of them.
package goalgo

import (
    "math/rand"
)

// In order to use goalgo.Shuffle(), you must first adapt your collection to
// this interface. Note that if you collection is sort.Sortable, it is already
// Shufflable.
type Shufflable interface {
    Size()			// number of elements in the collection
    Exch(i, j int)  // swap two elements at the given indices
}

// O(N). Randomly shuffle the given collection with Knuth Shuffle.
func Shuffle(a Shufflable) {
    for i := 0; i < a.Size(); i++ {
        j := i + rand.Intn(a.Size() - i)
        a.Exch(i, j)
    }
}
