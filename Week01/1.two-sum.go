/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	resultMap := make(map[int]int)
	for i, m := range nums {
		a := target - m
		if j, ok := resultMap[a]; ok {
			return []int{i, j}
		}
		resultMap[m] = i
	}
	return nil
}

// @lc code=end

