package lesson_25_testing

import (
	"testing"
	"reflect"
)

// Test explodeNonEmpty with empty string
func TestExplodeNonEmptyWithEmptyString(t *testing.T) {
	var actual []string
	testData := ""
	actual = explodeNonEmpty(testData)
	if len(actual) > 0 {
		t.Error("Not empty", actual)
	}
}

// Test explodeNonEmpty with five spaces
func TestExplodeNonEmptyWith5Spaces(t *testing.T) {
	var actual []string
	testData := "     "
	actual = explodeNonEmpty(testData)
	if len(actual) > 0 {
		t.Error("Not empty", actual)
	}
}

// Test explodeNonEmpty with two separated words
func TestExplodeNonEmptyWith2Words(t *testing.T) {
	testData := "first second"
	expected := []string{"first", "second"}
	actual := explodeNonEmpty(testData)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected = %v\nactual = %v", expected, actual)
	}
}

// Test explodeNonEmpty with many words and many spaces
func TestExplodeNonEmptyWithManyWordsAndManySpaces(t *testing.T) {
	testData := " first  second   third    fourth     fifth      "
	expected := []string{"first", "second", "third", "fourth", "fifth"}
	actual := explodeNonEmpty(testData)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected = %v\nactual = %v", expected, actual)
	}
}
