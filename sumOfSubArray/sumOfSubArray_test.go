package sumOfSubArray

import "testing"

func Test_rangeSum(t *testing.T) {
	actual := rangeSum([]int{1,2,3,4}, 4, 1, 5) 
	expected := 13
	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}
}