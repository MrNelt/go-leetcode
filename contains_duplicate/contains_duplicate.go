package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if set[v] {
			return true
		}
		set[v] = true
	}
	return false
}
