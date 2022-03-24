package main

import "testing"

func TestSum(t *testing.T) {
	a, b := 2, 3
	if sum(a, b) != 5 {
		t.Errorf("Sum(%d, %d) = %d, want %d", a, b, sum(a, b), 3)
	}
}

func TestProcessArgs(t *testing.T) {
	a, b := "2", "3"
	result, err := processArgs([]string{"sum", a, b})
	if err != nil {
		t.Error(err)
	}
	if result != "2 + 3 = 5" {
		t.Errorf("processArgs(%s, %s) = %s, want %s", a, b, result, "2 + 3 = 5")
	}
}
