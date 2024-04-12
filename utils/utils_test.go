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
