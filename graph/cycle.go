/*
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
*/

package graph

type cycleState struct {
    g *G
    seen, done []bool
}

func cycleDFS(state *cycleState, u int) bool {
    state.seen[u] = true
    for _, edge := range state.g.Adj(u) {
        v := edge.To()
        if !state.seen[v] {
            if cycleDFS(state, v) {
                return true
            }
        } else {
            if !state.done[v] {
                return true
            }
        }
    }
    state.done[u] = true
    return false
}

func HasCycle(g *G) bool {
    seen := make([]bool, g.V())
    done := make([]bool, g.V())
    state := &cycleState { g, seen, done }
    for u := 0; u < g.V(); u++ {
        if cycleDFS(state, u) {
            return true
        }
    }
    return false
}
