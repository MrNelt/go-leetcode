package two_sum

import (
	"sort"
)

func TwoSum(nums []int, target int) []int {

	data := make([][2]int, len(nums))
	for i, v := range nums {
		data[i][0] = v
		data[i][1] = i
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})

	l := 0
	r := len(data) - 1
	for l < r {
		sum := data[l][0] + data[r][0]
		if sum == target {
			return []int{data[l][1], data[r][1]}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{0, 0}
}
