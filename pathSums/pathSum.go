package pathSums

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    return dfs(root, targetSum, 0)
}
    

func dfs(root *TreeNode, targetSum int, pathSum int) bool {

    if root == nil{
        return false
    }

    pathSum+=root.Val
  
    if isLeaf(root, targetSum, pathSum){
        return true
    }

    return dfs(root.Left, targetSum, pathSum) || dfs(root.Right, targetSum, pathSum)
}    

func isLeaf(root *TreeNode, targetSum int, pathSum int) bool{
    if root.Right == nil && root.Left == nil{
        if targetSum == pathSum{
            return true
        }
    }
    return false
}