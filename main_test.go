package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	result := add(1, 4)
	expect := 5
	if result != expect {
		t.Errorf("Expected %v , but TestPlus result is %v.", expect, result)
	}
}

func TestMinus(t *testing.T) {
	result := minus(7, 3)
	expect := 4
	if result != expect {
		t.Errorf("Expected %v, but TestMinus result is %v.", expect, result)
	}
}

func TestProd(t *testing.T) {
	result := prod(8, 6)
	expect := 48
	if result != expect {
		t.Errorf("Expected %v, but TestProd result is %v", expect, result)
	}
}
