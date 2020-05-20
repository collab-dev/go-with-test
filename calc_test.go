package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(2, 2)
	if total != 4 {
		t.Error("Sum was incorrect")
	}
}

func TestAddMore(t *testing.T) {
	total := add(3, 3)
	if total != 6 {
		t.Error("Sum was incorrect")
	}
}
