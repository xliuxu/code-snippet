/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		val := digits[i] + 1
		if val < 10 {
			digits[i] = val
			return digits
		}
		digits[i] = val - 10
	}
	digits = append([]int{1}, digits...)
	return digits
}

