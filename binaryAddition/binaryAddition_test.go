package binaryAddition

import "testing"

func Test_addBinary(t *testing.T) {
	actual:= addBinary("1010", "1011")
	expected:="10101"

	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}
}