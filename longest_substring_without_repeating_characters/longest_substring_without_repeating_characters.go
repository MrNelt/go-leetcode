package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	ans := 0
	set := make(map[byte]bool)
	l := 0
	r := 0
	for r < len(s) {
		for set[s[r]] {
			set[s[l]] = false
			l++
		}
		set[s[r]] = true
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
