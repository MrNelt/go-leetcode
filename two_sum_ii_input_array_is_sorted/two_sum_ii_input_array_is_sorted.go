package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l, r}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{0, 0}
}
