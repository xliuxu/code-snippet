import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	res = append(res, []int{})
	for i := 0; i < len(nums); {
		count := 0
		for count+i < len(nums) && nums[i] == nums[count+i] {
			count++
		}
		for _, p := range res {
			cur := append(p[:0:0], p...)
			for j := 0; j < count; j++ {
				cur = append(cur, nums[i])
				res = append(res, cur)
			}
		}
		i += count
	}
	return res
}
