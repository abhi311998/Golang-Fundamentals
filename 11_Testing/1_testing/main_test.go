package main

import "testing"

func Test_mySum(t *testing.T) {
	x := mySum(2, 3, 4)
	if x != 9 {
		t.Error("Expected ", 9, " , Got", x)
	}
}
