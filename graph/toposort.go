/*  
    The topological order of a directed graph is a way to arrange all the
    vertices so that, when being put on a vertical straight line from the
    bottom to the top, all edges will point upwards. 

    Topological order is often used in task scheduling. You have a number of
    tasks where one depends on another and you want to find out in which order
    you should handle them. Software package dependency comes to mind here.

    For a topological order to exist, said directed graph must have no cycles.
    If we perform TopoSort() on a directed cyclic graph (DAG for short), the
    result can still be of interest to you (for example, in Kosaraju-Sharir's
    algorithm), but it isn't technically a topological order.

    Conveniently, if we do a depth first search and sort the vertices by the
    completion time of such search (vertices that completes earlier appears
    later), what we have will be a valid topological order.

    The reason why this works, and also why DFS tends to yield interesting
    consequences, is because DFS naturally forms a spanning forest. DFS also
    guarantees that it examines every edge, although it may not follow through.
    When DFS sees an edge u->v, there are three possible cases:

    (a) u is an ancestor of v in the DFS forest: u->v is called a forward edge.
    (b) v is an ancestor of u in the DFS forest: u->v is called a backward one.
    (c) otherwise, u and v belong to separate branches, and we shall call u->v
    a cross edge. We also know that every cross edge must point to a vertice
    explored earlier.

    The topological sorting algorithm described above basically puts the leaves
    at the end of the list, walks up the trees, and the roots will appear at 
    the beginning of the list.
    
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

    For an example usage, see toposort_example.go.
*/

package graph

type topoState struct {
    g *G
    a []int
    b []bool
    k int
}

func topoDFS(s *topoState, u int) {
    s.b[u] = true
    for _, e := range s.g.Adj(u) {
        v := e.To()
        if !s.b[v] {
            topoDFS(s, v)
        }
    }
    s.a[s.k] = u
    s.k--
}

// This function assumes that the given graph is acyclic. Call 
// graph.IsAcyclic() beforehand if you are unsure.

func TopoSort(g *G) []int {
    a := make([]int, g.V())
    b := make([]bool, g.V())
    k := g.V() - 1
    s := &topoState { g, a, b, k }
    for u := 0; u < g.V(); u++ {
        if !b[u] {
            topoDFS(s, u)
        }
    }
    return a
}