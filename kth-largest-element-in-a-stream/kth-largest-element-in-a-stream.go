package kthlargestelementinastream

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
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

type KthLargest struct {
	k    int
	heap IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	temp := IntHeap(nums)
	heap.Init(&temp)
	for len(temp) > k {
		heap.Pop(&temp)
	}
	return KthLargest{k: k, heap: temp}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.heap, val)
	if len(this.heap) > this.k {
		heap.Pop(&this.heap)
	}
	return this.heap[0]
}
