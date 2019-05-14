package leetcode

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	var res int
	for x != 0 {
		res *= 10
		res += x % 10
		x = x / 10
	}
	if res < -(1<<31) || res > (1<<31)-1 {
		return 0
	}
	return res
}
