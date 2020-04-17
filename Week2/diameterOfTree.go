/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var maxPath int = 0
	traverse(root, &maxPath)
	return maxPath
}

func traverse(root *TreeNode, maxPath *int) int {
	if root == nil {
		return 0
	} else {
		L := traverse(root.Left, maxPath)
		R := traverse(root.Right, maxPath)
		// fmt.Println("Node", root.Val, L + R + 1, *maxPath, Max(*maxPath, L + R + 1))
		*maxPath = Max(*maxPath, L+R)
		// fmt.Println("Node After Assign", L + R + 1, *maxPath)
		return Max(L, R) + 1
	}
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}