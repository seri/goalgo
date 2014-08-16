// QuickSort has quadaratic performance in the worst case and uses 1.39NlgN 
// compares on average. Because the worst case is highly unlikely thanks to the
// random shuffle which effectively acts as a performance shield, quicksort is
// faster than MergeSort in practice. We provide here two versions of QuickSort,
// the widely known one and Djikstra's QuickSort. The Djikstra version is
// preferrable when we know that the collection contains many duplicates. Oh and
// by the way, QuickSort is unstable.

package sort

import (
    "github.com/seri/goalgo"
)

func partition(a Sortable, lo, hi int) int {
    i, j := lo + 1, hi
    for {
        for ; i <= hi && a.Less(i, lo); i++ {}
        for ; a.Less(lo, j); j-- {}
        if i >= j {
            break
        }
        a.Exch(i, j)
        i, j = i + 1, j - 1
    }
    a.Exch(lo, j)
    return j
}

func quickSort(a Sortable, lo, hi int) {
    if hi <= lo {
        return
    }
    mid := partition(a, lo, hi)
    quickSort(a, lo, mid - 1)
    quickSort(a, mid + 1, hi)
}

func QuickSort(a Sortable) {
    goalgo.Shuffle(a)
    quickSort(a, 0, a.Size() - 1)
}

func djikstraPartition(a Sortable, lo, hi int) (lt, gt int) {
    // We maintain that a[(lo + 1)..(lt - 1)] < v = a[lt..gt] < a[(gt + 1)..hi]
    lt, gt = lo + 1, hi
    for i := lt; i <= gt; {
        switch {
        case a.Less(i, lo):
            a.Exch(i, lt)
            lt, i = lt + 1, i + 1
        case a.Less(lo, i):
            a.Exch(i, gt)
            gt = gt - 1
        default:
            i++
        }
    }
    lt--
    a.Exch(lo, lt)
    return
}

func djikstraQuickSort(a Sortable, lo, hi int) {
    if hi <= lo {
        return
    }
    lt, gt := djikstraPartition(a, lo, hi)
    djikstraQuickSort(a, lo, lt - 1)
    djikstraQuickSort(a, gt + 1, hi)
}

func DjikstraQuickSort(a Sortable) {
    goalgo.Shuffle(a)
    djikstraQuickSort(a, 0, a.Size() - 1)
}
