/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	know := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lengthOfString := len(s)
	lastElement := s[len(s)-1 : lengthOfString]
	var result int
	result = know[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		if know[s[i:i+1]] <= know[s[i-1:i]] {
			result += know[s[i-1:i]]
		} else {
			result -= know[s[i-1:i]]
		}
	}
	return result
}

// @lc code=end
