package main

import (
	"testing"
	"runtime"
)

func TestGetOSType(t *testing.T) {

	expected := runtime.GOOS
	actual := GetOSType()

	if expected != actual {
		t.Errorf("%q != %q", actual, expected)
	}
}
