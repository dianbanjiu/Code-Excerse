package leetcode

// 104. Maximum Depth of Binary Tree
func maxDepth(root *TreeNode) int {
	if root == nil { // If this is a empty tree, return 0
		return 0
	}

	var (
		i int        // index
		p []TreeNode // current level elements
		q []TreeNode // next level elements
	)
	p = append(p, *root)
	for len(p) != 0 {
		for _, v := range p {
			if v.Left != nil {
				q = append(q, *v.Left)
			}
			if v.Right != nil {
				q = append(q, *v.Right)
			}
		}
		p = q
		q = q[:0]
		i++
	}
	return i
}
