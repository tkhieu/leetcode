import "fmt"

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // Handle cases where k > n

	// Helper function to reverse a portion of the array
	reverse := func(arr []int, start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	// Reverse the entire array
	reverse(nums, 0, n-1)
	// Reverse the first k elements
	reverse(nums, 0, k-1)
	// Reverse the remaining n-k elements
	reverse(nums, k, n-1)
}

// @lc code=end
