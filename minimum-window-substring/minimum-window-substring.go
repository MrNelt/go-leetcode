package minimumwindowsubstring

import (
	"math"
)

func MinWindow(s string, t string) string {
	l := 0
	r := 0
	lmin := 0
	rmin := 0
	ans := math.MaxInt
	set1 := make(map[byte]int)
	set2 := make(map[byte]int)
	for i := range t {
		set1[t[i]]++
	}
	for r < len(s) {
		set2[s[r]]++
		for contain(set1, set2) {
			if r-l < ans {
				rmin = r
				lmin = l
				ans = r - l
			}
			set2[s[l]]--
			l++

		}
		r++
	}
	if ans != math.MaxInt {
		return s[lmin : rmin+1]
	}
	return ""
}

func contain(set1, set2 map[byte]int) bool {
	for i := range set1 {
		if set1[i] > set2[i] {
			return false
		}
	}
	return true
}
