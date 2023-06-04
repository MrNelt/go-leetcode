package laststoneweight

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	temp := IntHeap(stones)
	heap.Init(&temp)
	for temp.Len() > 1 {
		y := heap.Pop(&temp).(int)
		x := heap.Pop(&temp).(int)
		if y > x {
			heap.Push(&temp, y-x)
		}
	}
	if temp.Len() == 1 {
		return temp[0]
	}
	return 0
}
