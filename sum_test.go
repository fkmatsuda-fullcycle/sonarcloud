package main

import "testing"

func TestSum(t *testing.T) {
	a, b := 2, 3
	if sum(a, b) != 5 {
		t.Errorf("Sum(%d, %d) = %d, want %d", a, b, sum(a, b), 3)
	}
}
