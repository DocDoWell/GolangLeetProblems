package kdistinct

import "testing"

func TestMaxDepth(t *testing.T) {

	actual := kthDistinct([]string{"d","b","c","b","c","a"}, 2) 
	expected := "a"
	if actual != expected {
		t.Fatalf(`expecting %s but got %s`, expected, actual)
	}


}