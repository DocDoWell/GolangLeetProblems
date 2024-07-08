package searchMatrix

import (
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	firstInput:= [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	firstActual:= searchMatrix(firstInput, 3)
	firstExpected := true
	
	if firstActual != firstExpected{
		t.Fatalf(`expecting %v but got %v`, firstExpected, firstActual)
	}

	secondInput:= [][]int{{1}}
	secondActual:= searchMatrix(secondInput, 0)
	secondExpected := false

	if secondActual != secondExpected{
		t.Fatalf(`expecting %v but got %v`, secondExpected, secondActual)
	}
}