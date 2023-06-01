package permutation_in_string

func checkInclusion(s1 string, s2 string) bool {

	set := [26]int{}
	setTemp := [26]int{}

	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		set[s1[i]-'a']++
	}

	for i := 0; i < len(s1); i++ {
		setTemp[s2[i]-'a']++
	}

	if equal(setTemp, set) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		setTemp[s2[i-len(s1)]-'a']--
		setTemp[s2[i]-'a']++
		if equal(setTemp, set) {
			return true
		}
	}

	return false
}

func equal(m1, m2 [26]int) bool {
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}
