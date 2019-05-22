/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
func myAtoi(str string) int {
	var res int64
	var started, neg bool
	sb := []byte(str)
	for _, c := range sb {
		if !started {
			if isSpace(c) {
				continue
			} else if isSign(c) {
				started = true
				if c == '-' {
					neg = true
				}
				continue
			} else if isDigit(c) {
				started = true
				res = int64(c - '0')
				continue
			} else {
				break
			}
		} else {
			if isDigit(c) {
				if neg {
					res = res*10 - int64(c-'0')
				} else {
					res = res*10 + int64(c-'0')
				}
				if res < -(1 << 31) {
					return -(1 << 31)
				} else if res > (1<<31)-1 {
					return (1 << 31) - 1
				} else {
					continue
				}
			} else {
				break
			}
		}
	} // for
	return int(res)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSign(b byte) bool {
	return b == '+' || b == '-'
}

func isSpace(b byte) bool {
	return b == ' '
}
