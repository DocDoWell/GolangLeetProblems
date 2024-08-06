package minimumpushes

import "testing"

func TestMinimumPushes(t *testing.T) {

	actual := minimumPushes("aabbccddeeffgghhiiiiii") 
	expected := 24
	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}


}