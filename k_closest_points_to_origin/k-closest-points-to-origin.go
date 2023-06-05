package kclosestpointstoorigin

import "container/heap"

type PointsHeap [][]int

func (h PointsHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}

func (h PointsHeap) Len() int {
	return len(h)
}

func (h PointsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointsHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *PointsHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	temp := PointsHeap(points)
	heap.Init(&temp)
	for len(temp) > k {
		heap.Pop(&temp)
	}
	return temp
}
