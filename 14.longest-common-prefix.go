/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}
	return p
}

// @lc code=end

