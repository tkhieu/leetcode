/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	var a int
	var temp int = x
	if temp < 0 {
		return false
	}
	for true {
		a = a*10 + temp%10
		temp = temp / 10

		if temp == 0 {
			break
		}
	}
	return a == x
}

// @lc code=end
