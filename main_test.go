package main

import "testing"

func TestGetGreeting(t *testing.T) {
	g := getGreeting()

	if g != "Hello world!" {
		t.Fail()
	}
}