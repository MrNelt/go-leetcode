package sum3

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := [][]int{}
	l := 0
	set := make(map[[3]int]bool)
	for l < len(nums)-2 && nums[l] < 1 {
		l0 := l + 1
		r := len(nums) - 1
		for l0 < r {
			target := nums[l] + nums[l0] + nums[r]
			if target == 0 {
				v := [3]int{nums[l], nums[l0], nums[r]}
				if !set[v] {
					ans = append(ans, v[:3])
				}
				set[v] = true
				l0++
			} else if target > 0 {
				r--
			} else {
				l0++
			}
		}
		l++
	}
	return ans
}
