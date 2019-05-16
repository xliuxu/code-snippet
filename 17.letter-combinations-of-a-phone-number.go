
/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
func letterCombinations(digits string) []string {
	var ret []string
	if len(digits) == 0 {
		return ret
	}
	ret = append(ret, "")
	dbytes := []byte(digits)
	ar := []string{
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	for _, d := range dbytes {
		var appended []string
		for _, c := range ar[(d - '1')] {
			for _, r := range ret {
				appended = append(appended, r+string(c))
			}
		}
		ret = appended
	}
	return ret
}
