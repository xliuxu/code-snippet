// package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // res -> index
	for idx, n := range nums {
		if i, ok := m[target-n]; ok {
			return []int{idx, i}
		} else {
			m[n] = idx
		}
	}
	// should not goes here
	return nil
}
