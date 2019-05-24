/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	sb := []byte(s)

	var ps, pe int
	ps = 0
	pe = len(s) - 1
	for ps < pe {
		if !isAlphanumeric(sb[ps]) {
			ps++
			continue
		}
		if !isAlphanumeric(sb[pe]) {
			pe--
			continue
		}
		if !equalIgnoreCase(sb[ps], sb[pe]) {
			return false
		} else {
			ps++
			pe--
		}
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func equalIgnoreCase(a, b byte) bool {
	if a == b {
		return true
	}
	if a >= 'a' && a <= 'z' {
		a = a + 'A' - 'a'
	}
	if b >= 'a' && b <= 'z' {
		b = b + 'A' - 'a'
	}
	return a == b
}
