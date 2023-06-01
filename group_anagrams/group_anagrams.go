package group_anagrams

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)
	for _, v := range strs {
		key := [26]int{}
		for _, j := range v {
			key[j-rune('a')]++
		}
		anagrams[key] = append(anagrams[key], v)
	}
	answer := [][]string{}
	for _, v := range anagrams {
		answer = append(answer, v)
	}
	return answer
}
