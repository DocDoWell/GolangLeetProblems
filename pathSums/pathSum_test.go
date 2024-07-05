package pathSums

import (
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	leaf2:= &TreeNode{2,nil,nil}
	leaf7:= &TreeNode{3,nil,nil}
	root:= &TreeNode{1,leaf2,leaf7}
	pass:= hasPathSum(root,4)

	if !pass{
		t.Fatalf(`expecting true but got false`)
	}
}

func Test_hasPathSum1(t *testing.T) {
	leaf2:= &TreeNode{2,nil,nil}
	leaf7:= &TreeNode{3,nil,nil}
	root:= &TreeNode{1,leaf2,leaf7}
	pass:= hasPathSum(root,5)

	if pass{
		t.Fatalf(`expecting true but got false`)
	}
}