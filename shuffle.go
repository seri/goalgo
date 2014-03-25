/* 
    The Knuth Shuffle algorithm randomly shuffles anything that implements the 
    Shufflable interface. It runs in linear time.

    Note that if something is sort.Sortable, it is also Shufflable.
*/

package goalgo

import (
    "math/rand"
)

type Shufflable interface {
    Size() int
    Exch(i, j int)
}

func Shuffle(a Shufflable) {
    for i := 0; i < a.Size(); i++ {
        j := i + rand.Intn(a.Size() - i)
        a.Exch(i, j)
    }
}
