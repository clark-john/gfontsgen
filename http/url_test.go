package http_test

import (
	"testing"

	"github.com/clark-john/gfontsgen/http"
)

func TestSetUrlQuery(t *testing.T) {
	url := "http://yesno.wtf/api"
	expected := "http://yesno.wtf/api?strict=true"
	http.SetUrlQuery(&url, "strict", "true")
	if url != expected {
		t.Fatalf("expected %s received %s", expected, url)
	}
}
