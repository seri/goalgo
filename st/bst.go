package st

import (
    "container/list"
)

type bstNode struct {
    key Comparable
    value interface{}
    left, right *bstNode
}

// A BST is a naive binary search tree which suffers linear complexity in the worst
// case although it is actually logarithmic on average, which is certainly not too
// bad at all.
type BST struct {
    root *bstNode
}

func NewBST() *BST {
    return &BST { nil }
}

// O(1)
func (me BST) Empty() bool {
    return me.root == nil
}

// O(N) worst, O(lgN) average
func (me *BST) Contains(k Comparable) bool {
    return bstGet(me.root, k) != nil
}

// O(N) worst, O(lgN) average
func (me *BST) Get(k Comparable) interface{} {
    if node := bstGet(me.root, k); node != nil {
        return node.value
    }
    return nil
}

// O(N) worst, O(lgN) average
func (me *BST) Put(k Comparable, v interface{}) {
    me.root = bstPut(me.root, k, v)
}

// O(NlgN)
func (me *BST) Flatten() []Item {
    a := make([]Item, 0)
    q := list.New()
    q.PushBack(me.root)
    for q.Len() != 0 {
        x := q.Remove(q.Front()).(*bstNode)
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

// O(N) worst, O(lgN) average
func (me *BST) Remove(k Comparable) {
    me.root = bstRemove(me.root, k)
}

// Get the node with key k in the tree rooted at root
func bstGet(root *bstNode, k Comparable) *bstNode {
    if root == nil {
        return nil
    }
    switch x := k.Compare(root.key); {
    case x < 0:
        return bstGet(root.left, k)
    case x > 0:
        return bstGet(root.right, k)
    }
    return root
}

// Put a new item in the tree rooted at root, return the new root
func bstPut(root *bstNode, k Comparable, v interface{}) *bstNode {
    if root == nil {
        return &bstNode { k, v, nil, nil }
    }
    switch x := k.Compare(root.key); {
    case x < 0:
        root.left = bstPut(root.left, k, v)
    case x > 0:
        root.right = bstPut(root.right, k, v)
    default:
        root.value = v
    }
    return root
}

// Remove the node with key k in the tree rooted at root, return the new root
func bstRemove(root *bstNode, k Comparable) *bstNode {
    if root == nil {
        return nil
    }
    switch x := k.Compare(root.key); {
    case x < 0:
        root.left = bstRemove(root.left, k)
        return root
    case x > 0:
        root.right = bstRemove(root.right, k)
        return root
    }
    return bstRemoveThis(root)
}

// Remove this particular root provided that root != nil, return the substituted root
func bstRemoveThis(root *bstNode) *bstNode {
    switch {
    case root.left == nil:
        return root.right
    case root.right == nil:
        return root.left
    }
    newRight, newRoot := bstRemoveMin(root.right)
    newRoot.left = root.left
    newRoot.right = newRight
    return newRoot
}

// Remove the node with minimum key in the tree rooted at root provided that root != nil
func bstRemoveMin(root *bstNode) (newRoot, minNode *bstNode) {
    if root.left == nil {
        newRoot, minNode = root.right, root
        return
    }
    newRoot = root
    root.left, minNode = bstRemoveMin(root.left)
    return
}
