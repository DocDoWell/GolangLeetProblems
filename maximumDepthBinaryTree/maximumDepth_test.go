package maximumDepth

import "testing"

func TestMaxDepth(t *testing.T) {
	leaf9:= &TreeNode{9,nil,nil}
	leaf15:= &TreeNode{15,nil,nil}
	leaf7:= &TreeNode{7,nil,nil}
	node20:= &TreeNode{20,leaf15,leaf7}
	root:= &TreeNode{20,leaf9,node20}
	
	level:= maxDepth(root)

	if level !=3{
		t.Fatalf(`expecting 3 but got %v`, level)
	}
}
