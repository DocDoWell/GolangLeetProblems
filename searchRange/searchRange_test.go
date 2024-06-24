package searchRange

import (
	"testing"
)

func Test_searchRange(t *testing.T) {
	firstInput:= []int{5,7,7,8,8,10}
	firstAcutal:= searchRange(firstInput, 8)
	firstExpected := []int{3,4}

	secondAcutal:= searchRange(firstInput, 6)
	secondExpected := []int{-1,-1}

	thirdInput:=[]int{1}
	thirdAcutal:= searchRange(thirdInput, 1)
	thirdExpected := []int{0,0}

	for i:=0; i < len(firstExpected);i++{
		if firstAcutal[i] != firstExpected[i]{
			t.Fatalf(`expecting %v but got %v`, firstExpected[i], firstAcutal[i])
		}
	}

	for i:=0; i < len(secondExpected);i++{
		if secondAcutal[i] != secondExpected[i]{
			t.Fatalf(`expecting %v but got %v`, secondExpected[i], secondAcutal[i])
		}
	}
	
	for i:=0; i < len(thirdExpected);i++{
		if thirdAcutal[i] != thirdExpected[i]{
			t.Fatalf(`expecting %v but got %v`, thirdExpected[i], thirdAcutal[i])
		}
	}
}
