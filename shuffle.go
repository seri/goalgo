// Package goalgo provides implementations of some fundamental algorithms, most of
// which come from the Algorithms course of the university of Princeton. There are
// runnable examples for most of these algorithms.
package goalgo

import (
    "math/rand"
)

// In order to use goalgo.Shuffle(), you must first adapt your collection to
// this interface. Note that if you collection is sort.Sortable, it is already
// Shufflable.
type Shufflable interface {
    Size() int      // number of elements in the collection
    Exch(i, j int)  // swap two elements at the given indices
}

// O(N). Randomly shuffle the given collection with Knuth Shuffle. See
// sort/quick.go for an example use.
func Shuffle(a Shufflable) {
    for i := 0; i < a.Size(); i++ {
        j := i + rand.Intn(a.Size() - i)
        a.Exch(i, j)
    }
}
