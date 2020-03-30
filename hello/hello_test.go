package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	name := "something"
	got := Hello(name)
	want := "Hello, " + name

	if got != want {
		t.Errorf("got %q  want %q", got, want)
	}
}
