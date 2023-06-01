package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	ans := 0
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	checked := make(map[int]bool)
	for v := range set {
		if checked[v] {
			continue
		}
		down := v
		up := v
		for set[down] {
			checked[down] = true
			down--
		}
		for set[up] {
			checked[up] = true
			up++
		}
		ans = max(ans, up-down-1)
	}
	return ans
}

func max(ans, i int) int {
	if ans > i {
		return ans
	}
	return i
}
