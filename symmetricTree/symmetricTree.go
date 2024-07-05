package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root.Left == nil && root.Right == nil {
		return true
	} else {

		if root.Left == nil {
			return false
		}

		if root.Right == nil {
			return false
		}

		if root.Left.Val != root.Right.Val {
			return false
		}
	}

	return solve(root.Left, root.Right)
}

func solve(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if (left != nil && right == nil) || (left == nil && right != nil) {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return solve(left.Left, right.Right) && solve(left.Right, right.Left)
}
