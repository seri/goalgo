The pq package introduces the Priority Queue API. As long as a collection has a
size, can compare any two elements, can swap any two elements, and knows how to
add and remove things in a first-in-first-out manner, it can also behave as a
priority queue. 

This particular implementation of priority queue is known as the maximum heap
data structure, which supports push and pop operations in logarithmic time.

We also implement HeapSort, which uses 2NlgN compares in the worst case but in
practice, QuickSort and even MergeSort are faster. HeapSort is unstable.

## Example

* [heap_exercise.go](https://github.com/seri/goalgo/blob/master/examples/heap_exercise.go)
shows how you may use this package in practice.

## References

[1] http://algs4.cs.princeton.edu/24pq/
