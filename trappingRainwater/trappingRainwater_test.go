package trappingRainWater

import (
	"testing"
)

func Test_trap(t *testing.T) {

	firstOutput := trap([]int{2,1,0,2})
	secondOutput := trap([]int{4,2})
	thirdOutput := trap([]int{2,0,2})
	fourthOutput := trap([]int{4,2,0,3,2,5})
	fifthOutput := trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})

	if firstOutput != 3{
		t.Fatalf(`expecting 3 but got %v`, firstOutput)
	}

	if secondOutput != 0{
		t.Fatalf(`expecting 0 but got %v`, secondOutput)
	}

	if thirdOutput != 2{
		t.Fatalf(`expecting 2 but got %v`, thirdOutput)
	}

	if fourthOutput != 9{
		t.Fatalf(`expecting 9 but got %v`, fourthOutput)
	}

	if fifthOutput != 6{
		t.Fatalf(`expecting 6 but got %v`, fifthOutput)
	}
}