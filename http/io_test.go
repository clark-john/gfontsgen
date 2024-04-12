package http_test

import (
	"strings"
	"testing"

	"github.com/clark-john/gfontsgen/http"
)

func TestReadBody(t *testing.T) {
	expected := "aerodynamics"
	r := strings.NewReader(expected)
	result := http.ReadBody(r)
	if expected != result {
		t.Fatalf("Expected %s receieved %s", expected, result)
	}
}
