/*
    The balanced binary search tree that we have here in the course is Robert
    Sedgewick's left-leaning red back tree, which is quite a topic of
    controversy. By introducing a very simple restriction that all red nodes
    must lean leaft, Sedgewick claims that his tree is cleaner to implement
    than a typical red black tree while suffering no performance penalty. The
    critics say the reverse is true. Personally, I like the left-leaning red
    black tree.
*/

package st

import (
    "container/list"
)

type llrbNode struct {
    key Comparable
    value interface{}
    left, right *llrbNode
    color bool // color of the connection between this node and its parent;
               // the value true means red while false means black
}

type LLRB struct {
    root *llrbNode
}

func NewLLRB() *LLRB {
    return &LLRB { nil }
}

// The code for all access queries are going to be exactly the same as they are
// in a standard binary search tree. While it is possible to eliminate this code
// repetition by introducing a common Node interface, the added complexity is 
// not worth it.

func (me LLRB) Empty() bool {
    return me.root == nil
}

func (me *LLRB) Contains(k Comparable) bool {
    return llrbGet(me.root, k) != nil
}

func (me *LLRB) Get(k Comparable) interface{} {
    if node := llrbGet(me.root, k); node != nil {
        return node.value
    }
    return nil
}

func llrbGet(root *llrbNode, k Comparable) *llrbNode {
    if root == nil {
        return nil
    }
    switch x := k.Compare(root.key); {
    case x < 0:
        return llrbGet(root.left, k)
    case x > 0:
        return llrbGet(root.right, k)
    }
    return root
}

func (me *LLRB) Flatten() []Item {
    a := make([]Item, 0)
    q := list.New()
    q.PushBack(me.root)
    for q.Len() != 0 {
        x := q.Remove(q.Front()).(*llrbNode)
        if x == nil {
            continue
        }
        item := Item { x.key, x.value }
        a = append(a, item)
        q.PushBack(x.left)
        q.PushBack(x.right)
    }
    return a
}

/*
    Local operations on nodes which you can combine to implement global mutators
    on the tree, namely Put() and Remove(). Pay attention to how these helpers
    pass the redness around.
*/

// A nil node is considered black.

func llrbIsRed(node *llrbNode) bool {
    if node == nil {
        return false
    }
    return node.color
}

// Usually called when the right subtree is red. Pass this redness to the left.
// Do read the paper for an illustation.

func llrbRotateLeft(root *llrbNode) *llrbNode {
    right := root.right
    root.right = right.left
    right.left = root
    right.color = root.color
    root.color = true // the left subtree is now red
    return right // the right child becomes the new root
}

// Usually called when the left subtree is red. Pass this redness to the right.
// Do read the paper for an illustation.

func llrbRotateRight(root *llrbNode) *llrbNode {
    left := root.left
    root.left = left.right
    left.right = root
    left.color = root.color
    root.color = true // the right subtree is now red
    return left // the left child becomes the new root
}

// Flip all the colors of root and its two subtrees. Called either to pass up
// the redness of both children to the parent (for insertion), or to pass the 
// redness from the parent down to both children (for deletion).

func llrbFlipColor(root *llrbNode) *llrbNode {
    root.color = !root.color
    root.left.color = !root.left.color
    root.right.color = !root.right.color
    return root
}

// Fix the current subtree to maintain that the right child must be black and
// there won't be two immediate consecutive reds on the left.

func llrbFixUp(root *llrbNode) *llrbNode {
    if llrbIsRed(root.right) && !llrbIsRed(root.left) {
        root = llrbRotateLeft(root)
    }
    if llrbIsRed(root.left) && llrbIsRed(root.left.left) {
        root = llrbRotateRight(root)
    }
    if llrbIsRed(root.left) && llrbIsRed(root.right) {
        root = llrbFlipColor(root)
    }
    return root
}

/*
    To insert a new key-value pair, we make up a new node and put it in the tree
    in the same way that we did with a standard binary search tree. We shall
    color it red. We then walk up the tree right upto the root, fixing whatever
    violations that have been introduced due to this new red node, and we will 
    do so with the help of the above operations.
*/

func (me *LLRB) Put(k Comparable, v interface{}) {
    me.root = llrbPut(me.root, k, v)
    me.root.color = false // the root is always black
}

func llrbPut(root *llrbNode, k Comparable, v interface{}) *llrbNode {
    if root == nil {
        return &llrbNode { k, v, nil, nil, true } // a new node is always red
    }

    switch x := k.Compare(root.key); {
    case x < 0:
        root.left = llrbPut(root.left, k, v)
    case x > 0:
        root.right = llrbPut(root.right, k, v)
    default:
        root.value = v
    }

    return llrbFixUp(root)
}

/*
    Deletion in LLRB is serious business so bear it with me.
    So, the idea is, uhm...
    Actually, I don't get LLRB deletion either.
    If you have a clear explanation to share, I would appreciate.
*/

// Assume both children are black. We would like to introduce redness in the
// left subtree. If possible, we will pass this redness from the right one.
// Read the paper for an illustration.

func llrbMoveRedLeft(root *llrbNode) *llrbNode {
    llrbFlipColor(root)
    if llrbIsRed(root.right.left) {
        root.right = llrbRotateRight(root.right)
        root = llrbRotateLeft(root)
        llrbFlipColor(root)
    }
    return root
}

// Assume both children are black. We would like to introduce redness in the
// right subtree. If possible, we will pass this redness from the left one.
// You should draw a tree here to visualise.

func llrbMoveRedRight(root *llrbNode) *llrbNode {
    llrbFlipColor(root)
    if llrbIsRed(root.left.left) {
        root = llrbRotateRight(root)
        llrbFlipColor(root)
    }
    return root
}

// Consider the warm-up problem of removing the left-most leaf in our LLRB. We
// don't want to remove a black node as that would make the black height of 
// the left subtree lower than that of the right subtree. So as we traverse 
// left, we keep looking for opportunities to introduce redness in the left
// subtree.

func llrbRemoveMin(root *llrbNode) *llrbNode {
    if root.left == nil {
        return nil
    }
    if (!llrbIsRed(root.left)) && (!llrbIsRed(root.left.left)) {
        root = llrbMoveRedLeft(root)
    }
    root.left = llrbRemoveMin(root.left)
    return llrbFixUp(root)
}

// Pick the leftmost node of the tree rooted at root

func llrbLeftMost(root *llrbNode) *llrbNode {
    for root.left != nil {
        root = root.left
    }
    return root
}

// Finally combine all the helpers.

func (me *LLRB) Remove(k Comparable) {
    // this extra work is done to simplify the code
    if  me.Contains(k) {
        me.root = llrbRemove(me.root, k)
        me.root.color = false
    }
}

func llrbRemove(root *llrbNode, k Comparable) *llrbNode {
    if k.Compare(root.key) < 0 {
        if !llrbIsRed(root.left) && !llrbIsRed(root.left.left) {
            root = llrbMoveRedLeft(root)
        }
        root.left = llrbRemove(root.left, k)
    } else {
        // immediately steal redness from the left if available
        if llrbIsRed(root.left) {
            root = llrbRotateRight(root)
        }
        // notice that root.left must be nil too when this condition holds
        if k.Compare(root.key) == 0 && root.right == nil {
            return nil
        }
        // notice that we are guaranteed root.right != nil at this point
        if !llrbIsRed(root.right) && !llrbIsRed(root.right.left) {
            root = llrbMoveRedRight(root)
        }
        if k.Compare(root.key) == 0 {
            // swap root with the leftmost node of the right subtree
            leftmost := llrbLeftMost(root.right)
            root.key = leftmost.key
            root.value = leftmost.value
            // which would then be removed
            root.right = llrbRemoveMin(root.right)
        } else {
            root.right = llrbRemove(root.right, k)
        }
    }
    return llrbFixUp(root)
}
