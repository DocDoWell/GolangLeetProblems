package minimumcommonelement

import "testing"

func Test_getCommon1(t *testing.T) {
	actual := getCommon([]int{1,2,3}, []int{2,4})
	expected := 2

	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}	
}

func Test_getCommon2(t *testing.T) {
	actual := getCommon([]int{1,2,3}, []int{4,8})
	expected := -1

	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}	
}