package pivotInt

import "testing"

func Test1(t *testing.T) {

	actual := pivotInteger(8) 
	expected := 6
	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}


}

func Test2(t *testing.T) {

	actual := pivotInteger(1) 
	expected := 1
	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}


}

func Test3(t *testing.T) {

	actual := pivotInteger(4) 
	expected := -1
	if actual != expected {
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}


}