
/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	return recurPermute(nums, [][]int{})
}

func recurPermute(nums []int, prev [][]int) [][]int {
	if len(nums) == 0 {
		return prev
	}
	var ret [][]int
	if len(prev) == 0 {
		return recurPermute(nums[1:], append(ret, []int{nums[0]}))
	} else {
		for l := 0; l < len(prev); l++ {
			arr := prev[l]
			for i := 0; i <= len(arr); i++ {
				n := make([]int, len(arr)+1)
				copy(n[:i], arr[:i])
				copy(n[i+1:], arr[i:])
				n[i] = nums[0]
				ret = append(ret, n)
			}
		}
	}
	return recurPermute(nums[1:], ret)
}

