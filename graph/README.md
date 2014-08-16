The graph package works with edge-weighted directed graphs, a flexible type of
graphs that can be adjusted to behave like non-weighted directed graphs or
undirected graphs. However, if you want vertice-weighted graphs, for example,
you will have to look somwhere else.

There are many ways one can represent edge-weighted directed graphs. For the
sake of simplicity, we shall provide only one such representation. The
adjacent list saves some space and performs reasonably well in graphs typically
found in practice barring the densest ones.

## Graph algorithms

### Topological sort

The topological order of a directed graph is a way to arrange all the vertices
so that, when being put on a vertical straight line from the bottom to the top,
all edges will point upwards. 

Topological order is often used in task scheduling. You have a number of tasks
where one depends on another and you want to find out in which order you should
handle them. Software package dependency comes to mind here.

For a topological order to exist, said directed graph must have no cycles. If
we perform a topological sort on a directed cyclic graph, the result may still
be of interest to you (for example, in Kosaraju-Sharir's algorithm), but it
isn't technically a topological order.

Conveniently, if we do a depth first search and sort the vertices by the
completion time of such search (vertices that completes earlier appears later),
what we have will be a valid topological order.

The reason why this works, and also why depth first search tends to yield
interesting consequences, is because DFS naturally forms a spanning forest. DFS
also guarantees that it examines every edge, although it may not follow through.
When DFS sees an edge u->v, there are three possible cases:

* (a) u is an ancestor of v in the DFS forest: `u->v` is called a forward edge.
* (b) v is an ancestor of u in the DFS forest: `u->v` is called a backward one.
* (c) otherwise, u and v belong to separate branches, and we shall call `u->v`
a cross edge. We also know that every cross edge must point to a vertice
explored earlier.

The topological sorting algorithm described above basically puts the leaves
at the bottom of the list, walks up the trees, and the roots will appear at 
the top of the list.

Apparently for every forward edge u->v, v will appear later in the list,
because the algorithm always put descendants before ancestors.

Because the graph is acyclic, there won't be any back edge to consider. 
Why? Suppose there is a back edge u->v, v will be an ancestor of u, so 
there must be a path from v to u, which means the edge u->v complete a
a cycle. But this contradicts the fact that the graph is acyclic.

Finally, we want to prove that for every cross edge u->v, v will appear
later in the list. Thanks to the property of cross edges, we know that 
DFS explores v before u. Because v and u belong to separate branches, DFS
must have completed v before u. The algorithm will therefore put v later
than u in the list.

Now that we have proved that every edge in the graph does point from a
vertice earlier in the list to one that appears later in the list, we can
conclude that the algorithm is indeed correct.

### Cycle detection

An acyclic directed graph (DAG) is one that contains no cycle. A cycle is a
directed path which ends at the same vertice that it starts with. We are
interested in detecing whether a graph has a cycle.

We will first prove that a graph has a cycle if and only if there is a back
edge. Luckily, we have already proved from topological sorting that a back
edge immediately yields to a cycle, so we are left with proving that if the
graph has a cycle, there must be a back edge, too.

Let us consider the vertice u in the cycle that DFS explores first. Because
all the other vertices in the cycle are unmarked by the time DFS(u) starts,
they will all be descandants of a subtree rooted at u, including the vertice
v that points to u in the cycle. By definition, v->u is a back edge.

We have reduced cycle detection to back edge detection. Recall that DFS does
examine every edge of the graph, although it may not follow through. We want
to know in which scenarios will an edge be a back edge.

Cycle detection may seem trivial. At least, it would seem like an easier
problem to solve than, say, finding a topological order. In reality, the
problem is non-trivial enough that the father of Python has a blog post about
it. [1]

Here is a solution I came up with. It is different from the one found in the
course. [2]

The algorithm is straightfoward: Do a depth-first search. Not only that we
mark vertices as soon as we see them as in standard DFS, we also mark
vertices as soon as we they are done. If DFS ever meets a vertice v such that
v was marked as seen but not marked as done, the algorithm returns true.  
After DFS finishes visiting every vertice, the algorithm returns false.

We shall prove the correctness of the above algorithm in two phases. First,
let us prove that if the algorithm returns true, there is indeed a cycle.
The algorithm only returns true when DFS(u) sees a vertice v adjacent to u
and DFS(v) is still running. By the same logic as in topological sort, if
DFS(v) is still running and we are in the middle of DFS(u), there must be
a path from v to u. But there is also an edge u->v. So there is a cycle from
v back to itself.

The second phase is harder. We have to prove that if the algorithm returns
false, the graph must have no cycle.

When DFS(u) sees a vertice v, there are two possible scenarios:

(a) If v was not marked as seen, DFS(v) will be executed immediately. When
DFS(v) is running its course, if it ever sees u, the algorithm will return
true (because u is being marked as seen but not done). Because the algorithm
returns false, DFS(v) has never seen u, which means there is no path from v
to u.

(a) If v was marked as seen, it must also have been marked as done
(otherwise the algorithm would have returned true). Since DFS(v) has already
completed, every vertice reachable from v must have been completed too (by
the conjecture proved earlier in topological sort). We deduce that u is not
reachable from v.

If for every edge u->v, there is no path from v back to u, we can conclude
that the graph has no cycle.

For an example usage, see cycle_example.go.

[1] http://neopythonic.blogspot.com/2009/01/detecting-cycles-in-directed-graph.html
[2] http://algs4.cs.princeton.edu/42directed/DirectedCycle.java.html

### Strongly connected components

A strongly connected component is a group of vertices in a directed graph
where there exists a path between any pair of vertices in the component.

The following function, graph.SCC(), divides all the vertices of the given
graph into different strongly connected components such that it is
impossible to add any vertice to each of these components.

Think of this operation as a way to compress the graph in a useful way. The
compressed version is called the kernel graph of the original graph. In the
kernel graph, each strongly connected component of the original graph is
treated as a single vertice. And if there exists an edge between any vertice
of one component to any vertice of another one in the original graph, we add
an edge in the kernel graph.

The kernel graph will be acyclic, otherwise we could have merged some
components into one. For this property, the kernel graph is also called the
kernel DAG.

We will implement Kosaraju-Sharir's algorithm to solve this problem. We
simply run two DFS. The first DFS runs on the reverse graph of the original
graph and records the completion order (as in topological sort, but the 
result may not technically be a topological order if the graph has a cycle).
The second DFS runs on the original graph but in that particular order we
just found. For any vertice v that DFS(u) happens to see, we record v to be
in the same strong component as u.

I find this algorithm pretty magical. Easy to implement, but quite tricky to
understand why it works. B.Heap's algorithm to find permutations comes to
mind here. Let's try proving it anyway.

Suppose we are in the second DFS and happen to be in the middle of DFS(u).
We now see an unmarked vertice v and immediately record v to be in the same
component as u. Because DFS(u) reaches v, there is certainly a path from u
to v. We want to prove that there must also be a path from v to u. We assume
that no, there is no path from v to u, and try to reach a contradiction from
here.

Because v wasn't marked when DFS(u) sees it, u must appear before v in the
completion order that the first DFS found. We also know that there is path
from u to v but no path from v back to u in the original graph, we deduce
that there is path from v to u but no path from u back to v in the reverse
graph.

Because u was completed before v in the first DFS, 

### Minimum spanning tree

### Shortest paths

## Examples

1. [toposort_example.go](https://github.com/seri/goalgo/blob/master/examples/toposort_example.go)

## References

1. http://algs4.cs.princeton.edu/41undirected/
2. http://algs4.cs.princeton.edu/42directed/
3. http://algs4.cs.princeton.edu/43mst/
4. http://algs4.cs.princeton.edu/44sp/
5. http://www.seas.gwu.edu/~simhaweb/alg/lectures/module7/module7.html
6. https://www.ics.uci.edu/~eppstein/161/960220.html
7. http://www.cs.nyu.edu/courses/summer04/G22.1170-001/6a-Graphs-More.pdf
