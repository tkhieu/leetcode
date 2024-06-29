/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	appeared := make(map[int]int)
	for _, v := range nums {
		appeared[v]++
		if appeared[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

// @lc code=end