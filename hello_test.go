package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello("Seungmin")
	want := "Hello, Seungmin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}