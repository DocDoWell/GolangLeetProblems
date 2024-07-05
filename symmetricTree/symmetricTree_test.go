package symmetrictree

import "testing"

func Test_isSymmetricTree(t *testing.T) {
	leaf3 := &TreeNode{3, nil, nil}
	leaf4 := &TreeNode{4, nil, nil}
	node2a := &TreeNode{2, leaf3, leaf4}
	node2b := &TreeNode{2, leaf4, leaf3}
	root := &TreeNode{1, node2a, node2b}

	symmetry := isSymmetric(root)

	if !symmetry  {
		t.Fatalf(`expecting true but got false`)
	}
}