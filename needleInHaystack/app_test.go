package needleinhaystack

import (
	"testing"
)

func Test_strStr(t *testing.T) {

	firstOutput := strStr("sadbutsad", "sad")

	secondOutput := strStr("leetcode", "leeto")

	if firstOutput != 0{
		t.Fatalf(`expecting 0 but got %v`, firstOutput)
	}

	if secondOutput != -1{
		t.Fatalf(`expecting -1 but got %v`, secondOutput)
	}

}