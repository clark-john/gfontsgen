package key_test

import (
	"os"
	"testing"

	"github.com/clark-john/gfontsgen/key"
)

func TestGetApiKey(t *testing.T) {
	val := "somesortofapikey"
	os.Setenv(key.KEY_NAME, val)
	if key.GetApiKey() != val {
		t.Fatalf("Getting API key from environment failed")
	}
	os.Unsetenv(key.KEY_NAME)
	if key.GetApiKey() != "" {
		t.Fatalf("Emptying API key from environment failed")
	}
}
