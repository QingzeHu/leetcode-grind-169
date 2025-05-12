/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	magazineCharCount := make([]int, 26)

	for _, char := range magazine {
		magazineCharCount[char-'a']++
	}

	for _, char := range ransomNote {
		charIndex := char - 'a'

		if magazineCharCount[charIndex] > 0 {
			magazineCharCount[charIndex]--
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

