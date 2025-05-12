/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	twoStepsBefore := 1
	oneStepBefore := 2
	currentWays := 0

	for i := 3; i <= n; i++ {
		currentWays = twoStepsBefore + oneStepBefore
		twoStepsBefore = oneStepBefore
		oneStepBefore = currentWays
	}
	return oneStepBefore
}

// @lc code=end

