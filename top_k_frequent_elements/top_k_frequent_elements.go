package top_k_frequent_elements

import "sort"

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v]++
	}
	data := [][2]int{}
	for i, v := range dict {
		data = append(data, [2]int{v, i})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] > data[j][0]
	})
	ans := []int{}
	for i := 0; i < min(k, len(data)); i++ {
		ans = append(ans, data[i][1])
	}
	return ans
}

func min(k, i int) int {
	if k > i {
		return i
	}
	return k
}
