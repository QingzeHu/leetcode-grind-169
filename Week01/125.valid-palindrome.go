/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	front := 0
	end := len(s) - 1
	for front < end {
		front = nextAlphanumeric(s, front, true)
		end = nextAlphanumeric(s, end, false)
		if front >= end || front == -1 || end == -1 {
			break
		}
		if toLower(s[front]) != toLower(s[end]) {
			return false
		}
		front++
		end--
	}
	return true
}

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func isLetter(c byte) bool {
	return ('0' <= c && c <= '9') || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func nextAlphanumeric(s string, index int, forward bool) int {
	for i := index; i >= 0 && i < len(s); {
		if isLetter(s[i]) {
			return i
		}
		if forward {
			i++
		} else {
			i--
		}
	}
	return -1 // 或其他边界值
}

// @lc code=end

