package st

import (
    "testing"
    "math/rand"
)

func getBlackHeight(tree *LLRB) int {
    ret := 0
    for node := tree.root; node != nil; node = node.left {
        if !node.color {
            ret++
        }
    }
    return ret
}

func checkNode(t *testing.T, height int, node *llrbNode,
                             depth int, parentColor bool) {
    if node == nil {
        if depth != height {
            t.Errorf("invalid LLRB: some leaves have different black heights")
        }
        return
    }
    if !node.color {
        depth++
    } else if parentColor {
        t.Errorf("invalid LLRB: two consecutive red nodes found on a downward path")
        return
    }
    if node.right != nil && node.right.color {
        t.Errorf("invalid LLRB: there is a red link that leans right")
        return
    }
    checkNode(t, height, node.left, depth, node.color)
    checkNode(t, height, node.right, depth, node.color)
}

func checkTree(t *testing.T, tree *LLRB) {
    height := getBlackHeight(tree)
    checkNode(t, height, tree.root, 0, false)
}

func doPut(tree *LLRB) {
    for i := 0; i < 100; i++ {
        j := rand.Int() % 1000
        tree.Put(Int(j), i)
    }
}

func TestPut(t *testing.T) {
    for i := 0; i < 5; i++ {
        tree := NewLLRB()
        doPut(tree)
        checkTree(t, tree)
    }
}

func doRemove(tree *LLRB) {
    for i := 0; i < 100; i++ {
        j := rand.Int() % 1000
        tree.Remove(Int(j))
    }
}

func TestRemove(t *testing.T) {
    for i := 0; i < 5; i++ {
        tree := NewLLRB()
        doPut(tree)
        doRemove(tree)
        checkTree(t, tree)
    }
}
