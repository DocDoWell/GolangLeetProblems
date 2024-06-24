package canjump

import "testing"


func TestCanJump(t *testing.T) {

	firstTest:= canJump([]int{2,3,1,1,4})

	if firstTest==false {
		t.Fatalf(`expecting true but got %v`, firstTest)
	}
}
