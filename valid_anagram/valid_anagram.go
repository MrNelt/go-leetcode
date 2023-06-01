package valid_anagram

func IsAnagram(s string, t string) bool {
	dict1 := make(map[byte]int)
	dict2 := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		dict1[s[i]]++
		dict2[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		if dict1[s[i]] != dict2[s[i]] {
			return false
		}
	}
	return true
}
