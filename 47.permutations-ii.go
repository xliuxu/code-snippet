/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
func permuteUnique(nums []int) [][]int {
	return permuteUniqueRecur(nums, [][]int{})
}

func permuteUniqueRecur(nums []int, prev [][]int) [][]int {
	if len(nums) == 0 {
		return prev
	}
	if len(prev) == 0 {
		return permuteUniqueRecur(nums[1:], [][]int{[]int{nums[0]}})
	}
	var res [][]int
	for _, p := range prev {
		for i := 0; i <= len(p); i++ {
			var r []int = make([]int, len(p)+1)
			copy(r[:i], p[:i])
			copy(r[i+1:], p[i:])
			r[i] = nums[0]
			res = append(res, r)
			if i < len(p) && p[i] == nums[0] {
				break
			}
		}
	}
	return permuteUniqueRecur(nums[1:], res)
}