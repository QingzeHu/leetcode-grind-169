/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	curr := root.Val
	pVal, qVal := p.Val, q.Val
	if pVal > curr && qVal > curr {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if pVal < curr && qVal < curr {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

// @lc code=end

