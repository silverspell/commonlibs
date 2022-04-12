package commonlibs

import "testing"

func TestMd5(t *testing.T) {
	in := "golang"
	want := "21cc28409729565fc1a4d2dd92db269f"
	got := Md5(in)
	if got != want {
		t.Fatalf("want %s got %s\n", want, got)
	}
}

func TestSubstring(t *testing.T) {
	in := "in a gadda da vida"
	want := "gadda"
	got := Substring(in, 5, 5)
	if got != want {
		t.Fatalf("want %s got %s\n", want, got)
	}
}
