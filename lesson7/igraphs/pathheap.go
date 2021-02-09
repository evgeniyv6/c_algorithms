package igraphs

type path struct {
	val int
	nodes []int
}

// min path from pathes
// see https://golang.org/pkg/container/heap/:
// A heap is a tree with the property that each node is the minimum-valued node in its subtree.
// IntHeap example
type pathHeap []path

func (h pathHeap) Len() int {return len(h)}
func (h pathHeap) Less(i, j int) bool {return h[i].val < h[j].val}
func (h pathHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *pathHeap) Push(item interface{}) {
	*h = append(*h, item.(path))
}

func (h *pathHeap) Pop() interface{} {
	tmp := *h
	l := len(*h)
	old := tmp[l-1]
	*h = tmp[:l-1]
	return old
}


