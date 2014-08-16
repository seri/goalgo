Package sort hosts a handful of sorting algorithms abstracted as `Sorter`s.

* SelectionSort is an unstable quadratic sorter using `sqr(N)/2` comparisons and
`N` exchanges.

* InsertionSort is a quadratic sorter. It is stable, fast for small input, and
performs well in partially sorted collections. The last point is especially
important as this is the reason why InsertionSort is the only quadratic sorter
that is still widely used in practice. On average, it uses `sqr(N)/4`
compares and exchanges.

* ShellSort is an interesting generalisation of InsertionSort whose performance
is, interestingly enough, still unknown. This particular implementation of
ShellSort uses Knuth's `3n + 1` gap sequence, which achieves `O(N^(3/2)` in
complexity.

* MergeSort is theoretically optimal with guaranteed `~NlgN` compares.
Depending on the merging strategy, MergeSort may be stable or unstable. Our
implementation produces a stable sort.
    
    > By the way, our implementation cannot be done in Go in a generic manner.
    > In other words, `sort.MergeSort()` can only sort `[]int`.

* QuickSort has quadaratic performance in the worst case and uses `1.39NlgN` 
compares on average. Because the worst case is highly unlikely thanks to the
random shuffle which effectively acts as a performance shield, QuickSort is
faster than MergeSort in practice. We provide here two versions of QuickSort,
the widely known one and Djikstra's QuickSort. The Djikstra version is
preferrable when we know that the collection contains many duplicates. Oh and
by the way, QuickSort is unstable.

Perhaps more interesting than these well-known algorithms are the way that the
Go language approaches polymorphism through a flexible interface semantics. One
can use type-alias to adapt a known type to a new interface. One can embed an
interface into a struct for inheritance.

## Example

* [sort_client.go](https://github.com/seri/goalgo/blob/master/examples/sort_client.go)
tests the correctness of these algorithms and benchmarks them.

## References

[1] http://algs4.cs.princeton.edu/21elementary/
[2] http://algs4.cs.princeton.edu/22mergesort/
[3] http://algs4.cs.princeton.edu/23quicksort/
