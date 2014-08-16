The st package provides the symbol table API, which is more commonly known as
dictionaries or maps, and two trees that implement it.

* We provide, first of all, a naive binary search tree which suffers linear
complexity in the worst case although it is actually logarithmic on average,
which is certainly not too bad at all. 

* The self-balancing search tree that we have here is Robert Sedgewick's
left-leaning red back tree, which is quite a topic of controversy [4]. By
introducing a very simple restriction that all red nodes must lean left,
Sedgewick claims that his tree is cleaner to implement than a typical red black
tree while suffering no performance penalty. The critics say the reverse is
true. Personally, I like the left-leaning red black tree.

In order to use each of these symbol tables, you must first adapt your key
type to the `Comparable` interface. When you get values out of a `ST`, you also
have to manually type-cast them. 

## Examples

* [st_client.go](https://github.com/seri/goalgo/blob/master/examples/st_client.go)
test and benchmarks the two trees.

## References

1. http://algs4.cs.princeton.edu/32bst/
2. http://algs4.cs.princeton.edu/33balanced/
3. http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf
4. [Left-Leaning Red-Black Trees Considered Harmful](http://www.read.seas.harvard.edu/~kohler/notes/llrb.html)
