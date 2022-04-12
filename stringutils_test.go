package commonlibs

import "testing"

func TestSubstring(t *testing.T) {
	in := "in a gadda da vida"
	want := "gadda"
	got := Substring(in, 5, 5)
	if got != want {
		t.Fatalf("want %s got %s\n", want, got)
	}
}

func TestSubstringInPlace(t *testing.T) {
	in := "in a gadda da vida"
	want := "gadda"
	SubstringInPlace(&in, 5, 5)
	if in != want {
		t.Fatalf("want %s got %s\n", want, in)
	}
}
