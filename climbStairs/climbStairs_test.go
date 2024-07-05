package climbstairs

import "testing"

func Test_climbStairs(t *testing.T) {
	actual:= climbStairs(4)
	expected:=5

	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}
}