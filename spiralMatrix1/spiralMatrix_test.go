package spiralMatrix

import (
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	firstInput:= [][]int{{1,2,3},{4,5,6},{7,8,9}}
	acutal:= spiralOrder(firstInput)
	expected := []int{1,2,3,6,9,8,7,4,5}
	for i:=0; i < len(expected);i++{
		if acutal[i] != expected[i]{
			t.Fatalf(`expecting %v but got %v`, expected[i], acutal[i])
		}
	}
}