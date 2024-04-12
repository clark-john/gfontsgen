package gen_test

import (
	"net/url"
	"testing"

	"github.com/clark-john/gfontsgen/gen"
	"github.com/clark-john/gfontsgen/json"
)

func TestSpaceToPlus(t *testing.T){
	sample := "My Font"
	result := gen.SpaceToPlus(sample)
	expected := "My+Font"
	if expected != result {
		t.Fatalf("want %s got %s", expected, result)
	}
}

func TestGenerateUrl(t *testing.T){
	m := make(map[string]string)

	m["300"] = ""
	m["300italic"] = ""

	_url, _ := gen.GenerateUrl(json.FontItem{
		Family: "My Font",
		Variants: []string{"300", "300italic"},
		Subsets: []string{"cyrillic", "latin"},
		Version: "",
		LastModified: "",
		Files: m,
		Category: "",
		Menu: "",
		Kind: "",
	}, &gen.GenerateUrlOptions{
		Variants: []string{"300", "300i"},
		Copy: false,
	})

	u, _ := url.Parse(_url)

	if u.Query().Get("display") != "swap" {
		t.Fatalf("Query params in a generated url doesn't contain \"display\"")
	}
	
	familyQWant := "My+Font:300,300i"

	// plus sign gets decoded into actual space 
	if gen.SpaceToPlus(u.Query().Get("family")) != familyQWant {
		t.Fatalf("want %s got %s", familyQWant, u.Query().Get("family"))
	}
}
