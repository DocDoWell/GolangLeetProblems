package maximumDepth

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, level int) int {
    if root == nil{
        return level
    }
    level+=1
    return maximum(dfs(root.Left,level), dfs(root.Right,level)) 
}    

func maximum(x int, y int) int {
	if x > y {
		return x
	}
	return y
}