package valid_palindrome

import "unicode"

func isPalindrome(s string) bool {
	t := []rune{}
	for _, v := range s {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			t = append(t, unicode.ToLower(v))
		}
	}
	for i := 0; i < len(t)/2; i++ {
		if t[i] != t[len(t)-1-i] {
			return false
		}
	}
	return true
}
