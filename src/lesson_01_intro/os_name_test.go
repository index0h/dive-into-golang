package lesson_01_intro

import (
	"runtime"
	"testing"
)

func TestGetOSType(t *testing.T) {

	expected := runtime.GOOS
	actual := GetOSType()

	if expected != actual {
		t.Errorf("%q != %q", actual, expected)
	}
}
