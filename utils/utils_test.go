package utils_test

import (
	"testing"

	"github.com/clark-john/gfontsgen/utils"
)

func TestCapitalize(t *testing.T) {
	result := utils.Capitalize("presto")
	expected := "Presto"

	if result != expected {
		t.Fatalf(`expected %s I received "%s"`, expected, result)
	}
}

func TestTitleCase(t *testing.T) {
	result := utils.ToTitleCase("react native")
	expected := "React Native"

	if result != expected {
		t.Fatalf(`expected "%s" I received "%s"`, expected, result)
	}
}

func TestStringToBytes(t *testing.T) {
	result := utils.StringToBytes("ABC")
	expected := []byte{65,66,67}

	for i := 0; i < len(expected); i++ {
		if !(expected[i] == result[i]) {
			t.Fatalf("Byte encoding failed: expected %d, received %d", expected[i], result[i])
		}
	}
}
