package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Seungmin")

	got := buffer.String()
	want := "Hello, Seungmin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
