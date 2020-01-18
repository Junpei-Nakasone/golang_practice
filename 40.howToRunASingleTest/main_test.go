package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 4)
	expected := 6
	if result != expected {
		t.Errorf("add() test returned an unexpected result: got %v want %v", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := sub(2, 4)
	expected := -2
	if result != expected {
		t.Errorf("sub() test returned an unexpected result: got %v want %v", result, expected)
	}
}
