package parser

import (
	"testing"
)

func TestParseModules(t *testing.T) {
	src := []byte(`module "example" { source = "./x" }`)
	mods, err := ParseModules("test.tf", src)
	if err != nil {
		t.Fatal(err)
	}
	if len(mods) != 1 {
		t.Fatalf("expected 1 module, got %d", len(mods))
	}
	if mods[0].Name != "example" {
		t.Errorf("expected name 'example', got %s", mods[0].Name)
	}
}