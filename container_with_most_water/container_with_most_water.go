package container_with_most_water

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for l < r {
		ans = max(min(height[l], height[r])*(r-l), ans)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
