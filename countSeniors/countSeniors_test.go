package countSeniors

import "testing"

func TestCountSeniors(t *testing.T) {

	input:= [] string {"7868190130M7522","5303914400F9211","9273338290F4010"}
	actual := countSeniors(input)
	expected := 2

	if actual != expected{
		t.Fatalf(`expecting %v but got %v`, expected, actual)
	}


}