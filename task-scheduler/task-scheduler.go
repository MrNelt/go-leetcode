package taskscheduler

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    byte
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	n := len(*pq)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value byte, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	count := make(map[byte]int)
	lastTime := [26]int{}
	for _, value := range tasks {
		count[value]++
	}
	pq := make(PriorityQueue, len(count))
	index := 0
	for i, count := range count {
		item := &Item{priority: count, value: i, index: index}
		pq[i] = item
		index++
	}

	heap.Init(&pq)
	time := 0
	for len(pq) != 0 {
		time++
		item := heap.Pop(&pq).(*Item)
		diff := n + 1
		if lastTime[item.value-'A'] != 0 {
			diff = time - lastTime[item.value-'A']
		}
		time += max(n-diff+1, 0)
		lastTime[item.value-'A'] = time
		item.priority -= 1
		if item.priority > 0 {
			heap.Push(&pq, item)
		}
		fmt.Println(string(item.value), time)
	}
	return time
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
