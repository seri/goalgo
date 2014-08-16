/*
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

*/

package graph
