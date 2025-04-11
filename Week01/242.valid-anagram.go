/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := [26]int{}
	for _, char := range s {
		counts[char-'a']++
	}
	for _, char := range t {
		counts[char-'a']--
		if counts[char-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

