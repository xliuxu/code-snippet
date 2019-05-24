/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */
func longestPalindrome(s string) int {
	m := make(map[byte]int)
	sb := []byte(s)
	for _, b := range sb {
		if _, ok := m[b]; ok {
			m[b] += 1
		} else {
			m[b] = 1
		}
	}
	var res int
	var hasOdd bool
	for _, v := range m {
		res += v / 2 * 2
		if v%2 != 0 {
			hasOdd = true
		}
	}
	if hasOdd {
		return res + 1
	} else {
		return res
	}
}

