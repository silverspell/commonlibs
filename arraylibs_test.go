package commonlibs

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}
	got := Filter(arr, func(item int) bool { return item%2 == 0 })
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v got %v", want, got)
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	want := []int{2, 4, 6, 8}
	got := Map(arr, func(item int) int { return item * 2 })
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v got %v", want, got)
	}
}

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	want := 10
	got := Reduce(arr, 0, func(k int, t int) int { return k + t })
	if got != want {
		t.Fatalf("want %v got %v", want, got)
	}
}

func TestFilterInPlace(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}
	FilterInPlace(&arr, func(item int) bool { return item%2 == 0 })
	if !reflect.DeepEqual(arr, want) {
		t.Fatalf("want %v got %v", want, arr)
	}
}

func TestMapInPlace(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	want := []int{2, 4, 6, 8}
	MapInPlace(&arr, func(item int) int { return item * 2 })
	if !reflect.DeepEqual(arr, want) {
		t.Fatalf("want %v got %v", want, arr)
	}
}

func TestStringReverse(t *testing.T) {
	in := "Hello World"
	chars := []rune(in)
	want := "dlroW olleH"
	Reverse(&chars)
	if string(chars) != want {
		t.Fatalf("want %v got %v", want, string(chars))
	}
}

func TestIntReverse(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	want := []int{5, 4, 3, 2, 1}
	Reverse(&in)
	if !reflect.DeepEqual(in, want) {
		t.Fatalf("want %v got %v", want, in)
	}
}
