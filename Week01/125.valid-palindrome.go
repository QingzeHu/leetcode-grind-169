/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	cleaned := []byte{}
	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			cleaned = append(cleaned, toLower(s[i]))
		}
	}

	front, end := 0, len(cleaned)-1
	for front < end {
		if cleaned[front] != cleaned[end] {
			return false
		}
		front++
		end--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

// @lc code=end

