6/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    startingColor := image[sr][sc]
	if startingColor == color {
        return image
    }

	rows, cols := len(image), len(image[0])

	var dfs func(r,c int)
	dfs = func(r,c int){
		if r < 0 || r >= rows || c < 0 || c >= cols || image[r][c] != startingColor {
            return
        }

		image[r][c] = color

		dfs(r+1,c)
		dfs(r-1,c)
		dfs(r,c+1)
		dfs(r,c-1)
	}

	dfs(sr,sc)
	return image
}
// @lc code=end

