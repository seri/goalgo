package main

import (
    "strconv"
    "strings"
    "github.com/seri/goalgo/st"
    . "./util"
)

func ToItemSlice(in string) []st.Item {
    a := ToIntSlice(in)
    b := make([]st.Item, len(a))
    for i := range a {
        b[i] = st.Item { st.Int(a[i]), true }
    }
    return b
}

func Stringify(t *st.BST) string {
    a := t.Flatten()
    b := make([]string, len(a))
    for i := range a {
        b[i] = strconv.Itoa(int(a[i].Key.(st.Int)))
    }
    return strings.Join(b, " ")
}

func ConstructBST(in string) *st.BST {
    t := st.NewBST()
    a := ToIntSlice(in)
    for _, x := range a {
        t.Put(st.Int(x), true)
    }
    return t
}

func BSTInsertion(in string) string {
    return Stringify(ConstructBST(in))
}

func HibbardDeletion(in string) string {
    a := strings.Split(in, ";")
    t := ConstructBST(a[0])
    for _, x := range ToIntSlice(a[1]) {
        t.Remove(st.Int(x))
    }
    return Stringify(t)
}

func main() {
    Test(BSTInsertion,
         "99 31 13 56 40 95 10 97 72 80",
         "99 31 13 56 10 40 95 72 97 80")
    Test(HibbardDeletion,
         "97 15 53 45 79 23 55 24 73 29 58 33;33 45 53",
         "97 15 55 23 79 24 73 29 58")
    Test(HibbardDeletion,
         "81 16 93 52 91 97 51 74 84 49 71 66; 84 16 52",
         "81 66 93 51 74 91 97 49 71")
    Test(BSTInsertion,
         "76 33 89 67 41 31 10 46 58 24",
         "76 33 89 31 67 10 41 24 46 58")
    Test(HibbardDeletion,
         "22 21 63 30 76 24 53 99 32 81 50 42;42 53 63",
         "22 21 76 30 99 24 32 81 50")
}
