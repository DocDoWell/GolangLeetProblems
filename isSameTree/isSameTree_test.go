package issametree
import "testing"

func Test_isSameTree(t *testing.T) {
	leaf9 := &TreeNode{9, nil, nil}
	leaf15 := &TreeNode{15, nil, nil}
	leaf7 := &TreeNode{7, nil, nil}
	node20 := &TreeNode{20, leaf15, leaf7}
	root := &TreeNode{20, leaf9, node20}

	same := isSameTree(root, root)
	different := isSameTree(root, leaf15)

	if !same  {
		t.Fatalf(`expecting true but got false`)
	}

	if different{
		t.Fatalf(`expecting false but got true`)
	}
}
