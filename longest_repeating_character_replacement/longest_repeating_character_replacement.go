package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	ans := 0
	for i := 0; i < 26; i++ {
		ans = max(ans, solve(s, byte('A'+i), k))
	}
	return ans
}

func solve(s string, b byte, k int) int {
	ans := 0
	l := 0
	r := 0
	temp := 0
	for r < len(s) {
		if s[r] != b {
			temp++
		}
		for temp > k {
			if s[l] != b {
				temp--
			}
			l++
		}
		r++
		ans = max(ans, r-l+1)

	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
