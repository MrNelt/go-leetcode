package sortedsquares

import "sort"

func SortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}
